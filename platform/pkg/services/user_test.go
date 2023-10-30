package services

import (
	"livecom/pkg/db"
	"livecom/pkg/repo"
	"testing"

	"firebase.google.com/go/v4/auth"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestMergeTokenUsrPayload(t *testing.T) {
	// Mock data
	token := &auth.Token{
		Claims: map[string]interface{}{
			"user_id":       "12345",
			"email":         "test@example.com",
			"email_verified": "true",
			"name":          "John Doe",
			"phone_number":  "1234567890",
			"photo":"https://photo.com/photo.jpg",
		
		},
	}

	tests := []struct {
		name     string
		token    *auth.Token
		payload  repo.CreateUserPayload
		expected *db.CreateUserParams
		hasError bool
	}{
		{
			name: "Valid token and Null payload",
			token: token,
			payload: repo.CreateUserPayload{
			
			},
			expected: &db.CreateUserParams{
				EmailVerified: pgtype.Bool{
					Bool: true,
					Valid: true,
				},
				FirebaseUid: pgtype.Text{
					String: "12345",
					Valid: true,
				},
				Email: "test@example.com",
				DisplayName: pgtype.Text{
					String: "John Doe",
					Valid: true,
				},
				FirstName: pgtype.Text{	},			
				LastName: pgtype.Text{	},
				PhoneNumber:pgtype.Text{
					String: "1234567890",
					Valid: true,
				},
				PhotoUrl: pgtype.Text{
					String: "https://photo.com/photo.jpg",
					Valid: true,
				},
			},
			hasError: false,
		},		
		{
			name: "Valid token and valid Payload",
			token: token,
			payload: repo.CreateUserPayload{
				FirstName: ptrToString("Joanna"),
				LastName: ptrToString("Dark"),
				PhoneNumber: ptrToString("1123321321"),
				PhotoUrl: ptrToString("https://example.com/photo.jpg"),
			
			},
			expected: &db.CreateUserParams{
				EmailVerified: pgtype.Bool{
					Bool: true,
					Valid: true,
				},
				FirebaseUid: pgtype.Text{
					String: "12345",
					Valid: true,
				},
				Email: "test@example.com",
				DisplayName: pgtype.Text{
					String: "John Doe",
					Valid: true,
				},
				FirstName: pgtype.Text{ String: "Joanna" ,Valid: true },			
				LastName: pgtype.Text{	String: "Dark" ,Valid: true},
				PhoneNumber:pgtype.Text{
					String: "1123321321",
					Valid: true,
				},
				PhotoUrl: pgtype.Text{
					String: "https://example.com/photo.jpg",
					Valid: true,
				},
			},
			hasError: false,
		},
		{
			name: "Valid token and Invalid Payload",
			token: token,
			payload: repo.CreateUserPayload{
				FirstName: ptrToString("Joanna"),
				LastName: ptrToString("Da"),
				PhoneNumber: ptrToString("1123321321"),
				PhotoUrl: ptrToString("https://example.com/photo.jpg"),
			
			},
			expected: &db.CreateUserParams{
				EmailVerified: pgtype.Bool{
					Bool: true,
					Valid: true,
				},
				FirebaseUid: pgtype.Text{
					String: "12345",
					Valid: true,
				},
				Email: "test@example.com",
				DisplayName: pgtype.Text{
					String: "John Doe",
					Valid: true,
				},
				FirstName: pgtype.Text{ String: "Joanna" ,Valid: true },			
				LastName: pgtype.Text{	String: "Da" ,Valid: true},
				PhoneNumber:pgtype.Text{
					String: "1123321321",
					Valid: true,
				},
				PhotoUrl: pgtype.Text{
					String: "https://example.com/photo.jpg",
					Valid: true,
				},
			},
			hasError: true,
		},
	}


	service := &Service{}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.mergeTokenUsrPayload(tt.token, tt.payload)

			if tt.hasError {
				assert.Error(t, err)
			} else {
					var uid pgtype.UUID;
					u := uuid.UUID(result.UserID.Bytes);
					uuidStr := u.String()
					uid.Scan(uuidStr);
				if err != nil {
					t.Logf("Error scanning UUID: %v", err)
					return
				}
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expected.FirstName, result.FirstName)
				assert.Equal(t, tt.expected.LastName, result.LastName)
				assert.Equal(t, tt.expected.Email, result.Email)
				assert.Equal(t, tt.expected.EmailVerified, result.EmailVerified)
				assert.Equal(t, tt.expected.DisplayName, result.DisplayName)
				assert.Equal(t, tt.expected.PhoneNumber, result.PhoneNumber)
				assert.Equal(t, tt.expected.PhotoUrl, result.PhotoUrl)
				assert.Equal(t, uid, result.UserID)
				assert.Equal(t, tt.expected.FirebaseUid, result.FirebaseUid)
			}
		})
	}
}
	
	// // Assert the results
	// var uid pgtype.UUID;
	// uid.Scan(result.UserID.Bytes);
	// assert.NoError(t, err)
	// assert.NotNil(t, result)
	// assert.Equal(t, "test@example.com", result.Email)
	// assert.Equal(t,"John Doe" , result.DisplayName.String)
	// assert.True(t, result.EmailVerified.Bool)
	//  assert.Equal(t, "John", result.FirstName.String)
	//  assert.Equal(t, "Doe", result.LastName.String)
	//  assert.Equal(t, "0987654321", result.PhoneNumber.String)
	//  assert.Equal(t, uid, result.UserID)
	//  assert.Equal(t, "http://example.com/photo.jpg", result.PhotoUrl.String)




//Mock Example
/* 
// Mock the repo
type MockRepo struct {
	mock.Mock
}


func (m *MockRepo) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
    args := m.Called(ctx,arg)
    return args.Get(0).(db.User), args.Error(1)
}

func (m *MockRepo) GetUserByID(ctx context.Context, userID pgtype.UUID) (db.User, error) {
    args := m.Called(ctx,userID)
    return args.Get(0).(db.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	// Initialize mock repo
	mockRepo := new(MockRepo)

	mockToken := &auth.Token{
		AuthTime: 1234567890,
		Issuer:   "mockIssuer",
		Audience: "mockAudience",
		Expires:  1234567890,
		IssuedAt: 1234567890,
		Subject:  "mockSubject",
		UID:      "mockUID",
		Firebase: auth.FirebaseInfo{
			// ... other fields ...
		},
		Claims: map[string]interface{}{
			"email": "test@example.com",
			"email_verified": true,
			"name":"Victor",
		},
	}

	// Test cases
	tests := []struct {
		name     string
		token    *auth.Token
		payload  repo.CreateUserPayload
		mockDbUserToCreate db.CreateUserParams
		createdUser db.User
		mockErr  error
	}{
		{
			name:    "Valid User Creation",
			token:   mockToken,
			payload: repo.CreateUserPayload{
				
			},
			mockDbUserToCreate: db.CreateUserParams{
				Email:       "test@example.com",
				EmailVerified: NewPGBool(true),
				// FirstName:   NewPGText("Jon"),
				DisplayName: NewPGText("Victor"),
				// LastName:    NewPGText("Doe"),
				// PhoneNumber: NewPGText("+1234567890"),
				// PhotoUrl:    NewPGText("https://example.com/photo.jpg"),
			
			},
			createdUser: db.User{
				Email:       "test@example.com",
				EmailVerified: NewPGBool(true),
				// FirstName:   NewPGText("Jon"),
				DisplayName: NewPGText("Victor"),
				// LastName:    NewPGText("Doe"),
				// PhoneNumber: NewPGText("+1234567890"),
				// PhotoUrl:    NewPGText("https://example.com/photo.jpg"),
			
			},
			mockErr:  nil,
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock expectations
			mockRepo.On("CreateUser",  mock.Anything, mock.MatchedBy(func(params db.CreateUserParams) bool {
				// Check the fields of params and return true if they match what you expect
				return true
			})).Return(tt.createdUser, tt.mockErr)
			// mockRepo.On("CreateUser", context.TODO(), tt.token, tt.payload).Return(tt.mockUser, tt.mockErr)
				s := &Service{
					db: mockRepo,
				}
		
			// // Call the function
			 user, err := s.CreateUser(context.TODO(), tt.token, tt.payload)
			 t.Logf("The value is: %v", user)
			 assert.IsType(t, pgtype.UUID{}, user.UserID)
			 assert.Equal(t, tt.createdUser.Email, user.Email)
			 assert.Equal(t, tt.createdUser.EmailVerified, user.EmailVerified)
			 assert.Equal(t, tt.createdUser.DisplayName, user.DisplayName)
			 // ... add other fields as needed ...
 
			 assert.Equal(t, tt.mockErr, err)
 
			 // Ensure expectations were met
			 mockRepo.AssertExpectations(t)
		})
	}
}
 */
// Helper function to get a pointer to a string
func ptrToString(s string) *string {
	return &s
}
func NewPGText(val string) pgtype.Text {
    var t pgtype.Text
    t.Scan(val)
    return t
}

func NewPGBool(val bool) pgtype.Bool {
    var t pgtype.Bool
    t.Scan(val)
    return t
}