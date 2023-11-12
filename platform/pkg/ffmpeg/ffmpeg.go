package ffmpeg

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"livecom/logger"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
	// Importar outras bibliotecas necessárias
)

type Resolution struct {
	Width    int
	Height   int
	BitRate  string
}
func (r Resolution) String() string {
	return fmt.Sprintf("{Width: %d, Height: %d, BitRate: %s}", r.Width, r.Height, r.BitRate)
}
type OutputFormat int

const (
    FormatNone OutputFormat = iota
    FormatDASH
    FormatHLS
    FormatBoth
)

type TranscodeParams struct {
	Resolutions []Resolution
	OutputType      OutputFormat
}
func (v *TranscodeParams) String() string {
	var resolutions []string
	for _, res := range v.Resolutions {
		resolutions = append(resolutions, res.String())
	}

	var outputTypeStr string
	switch v.OutputType {
	case FormatDASH:
		outputTypeStr = "DASH"
	case FormatHLS:
		outputTypeStr = "HLS"
	case FormatBoth:
		outputTypeStr = "Both DASH and HLS"
	default:
		outputTypeStr = "None"
	}

	return fmt.Sprintf("TranscodeParams {Resolutions: [%s], OutputType: %s}", strings.Join(resolutions, ", "), outputTypeStr)
}

type FFmpegTask struct {
	StreamUrl   string  `json:"stream_url"`  


	// FFmpeg pid.
	PID int32 `json:"pid"`
	// FFmpeg last frame.
	frame string
	// The last update time.
	update string

	// The context for current task.
	cancel context.CancelFunc


	transcode *TranscodeParams



	// To protect the fields.
	lock sync.Mutex
}

func (v *FFmpegTask) String() string {
	return fmt.Sprintf("StreamUrl=%v, PID=%v, , config is %v",
		v.StreamUrl, v.PID,  v.transcode.String(),
	)
}

func (v *FFmpegTask) saveTask(ctx context.Context) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	if _, err := json.Marshal(v); err != nil {
		return errors.Wrapf(err, "marshal %v", v.String())
	} 

	return nil
}


func (v *FFmpegTask) cleanup(ctx context.Context) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	if v.PID <= 0 {
		return nil
	}

	logger.Wf(ctx, "kill task pid=%v , stream_url=%v" , v.PID,v.StreamUrl)
	syscall.Kill(int(v.PID), syscall.SIGKILL)

	v.PID = 0
	v.cancel = nil

	return nil
}

func (v *FFmpegTask) Restart(ctx context.Context) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	if v.cancel != nil {
		v.cancel()
	}

	return nil
}
func (v *FFmpegTask) Start(ctx context.Context) error {
	// Create context for current task.
	parentCtx := ctx
	ctx, v.cancel = context.WithCancel(ctx)

}

// Transmissao representa uma transmissão individual e seu processo ffmpeg
type Transmissao struct {
	StreamUrl     string    // Identificador único para a transmissão
	Cmd    *exec.Cmd // Comando ffmpeg para esta transmissão
}

var (
	// Mutex para controlar o acesso ao mapa de transmissões
	mu sync.Mutex
	// Mapa para manter as transmissões ativas
	transmissoesAtivas = make(map[string]*Transmissao)
)

// IniciarTransmissao inicia uma nova transmissão
func IniciarTransmissao(ctx context.Context , id string) {


	mu.Lock()
	defer mu.Unlock()
	// Comando ffmpeg para esta transmissão
	cmd := exec.Command("ffmpeg", 
    "-f", "flv", 
    "-i", "rtmp://127.0.0.1:1935/live?vhost=__defaultVhost__/livestream", 
    "-map", "0:v", "-map", "0:a", 
    "-s:v:0", "640x360", "-aspect:v:0", "16:9", "-b:v:0", "600k", "-c:v:0", "libx264", "-c:a:0", "aac", "-b:a:0", "64k", 
    "-map", "0:v", "-map", "0:a", 
    "-s:v:1", "854x480", "-aspect:v:1", "16:9", "-b:v:1", "1200k", "-c:v:1", "libx264", "-c:a:1", "aac", "-b:a:1", "64k", 
    "-f", "dash", "-use_template", "1", "-use_timeline", "1", 
    "-adaptation_sets", "id=0,streams=v id=1,streams=a", 
    "-y", "./objs/nginx/html/live/my.mpd")

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		logger.E(ctx, "Error creating stdout pipe:", err)
		return
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		logger.E(ctx, "Error creating stderr pipe:", err)
		return
	}

	cmd.Stderr = os.Stderr
cmd.Stdout = os.Stdout

	// Start the ffmpeg process
	if err := cmd.Start(); err != nil {
		logger.E(ctx, "Error starting ffmpeg:", err)
		return
	}

	// Use separate goroutines to log stdout and stderr
	go streamToLogger(ctx, traceLogger, stdoutPipe)
	go streamToLogger(ctx, errorLogger, stderrPipe)



	// Armazenar a transmissão ativa
	transmissoesAtivas[id] = &Transmissao{
		StreamUrl:  id,
		Cmd: cmd,
	}
}

// FinalizarTransmissao termina uma transmissão ativa
func FinalizarTransmissao(c context.Context, id string) {
	mu.Lock()
	defer mu.Unlock()

	// Encontrar a transmissão pelo ID
	transmissao, ok := transmissoesAtivas[id]
	if !ok {
		logger.Ef(c,"Transmissão com ID %s não encontrada", id)
		return
	}

	// Terminar o processo ffmpeg
	if err := transmissao.Cmd.Process.Kill(); err != nil {
		logger.Ef(c,"Falha ao terminar o processo ffmpeg: ", err)
	
	}

	// Remover do mapa
	delete(transmissoesAtivas, id)
}



type LoggerWriter struct {
	LogFunc func(ctx logger.Context, a ...interface{})
	Ctx    logger.Context
}

func (lw *LoggerWriter) Write(p []byte) (n int, err error) {
	// Use the logging function (like logger.T, logger.Info, etc.) to log the message
	lw.LogFunc(lw.Ctx, string(p))
	return len(p), nil
}

func streamToLogger(ctx context.Context, logFunc func (ctx context.Context, a ...interface{}), r io.Reader) {
	flushInterval := 5 * time.Second // Flush every 5 seconds
	maxBufferSize := 3000          // Or any other appropriate buffer size
	var buffer bytes.Buffer
	flushTimer := time.NewTimer(flushInterval)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		buffer.WriteString(line + "\n")

		if buffer.Len() >= maxBufferSize {
			logFunc(ctx, buffer.String())
			buffer.Reset()
			flushTimer.Reset(flushInterval)
		}

		select {
		case <-flushTimer.C:
			if buffer.Len() > 0 {
				logFunc(ctx, buffer.String())
				buffer.Reset()
			}
			flushTimer.Reset(flushInterval)
		default:
			// Continue scanning
		}
	}

	// Handle any remaining data in buffer
	if buffer.Len() > 0 {
		logFunc(ctx, buffer.String())
	}

	if err := scanner.Err(); err != nil {
		logFunc(ctx, "Error reading from pipe:", err)
	}
}

// Wrapper for the logger.T function
func traceLogger(ctx context.Context, args ...interface{}) {
    // Assuming logger.Context can be directly converted from context.Context
    logger.T(ctx.(logger.Context), args...)
}

// Wrapper for the logger.E function
func errorLogger(ctx context.Context, args ...interface{}) {
    // Assuming logger.Context can be directly converted from context.Context
    logger.E(ctx.(logger.Context), args...)
}