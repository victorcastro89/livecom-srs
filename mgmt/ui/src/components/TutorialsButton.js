import React from "react";
import {Toast} from "react-bootstrap";
import logo from '../resources/logo.svg';
import * as Icon from 'react-bootstrap-icons';
import {Token} from "../utils";
import axios from "axios";

function useTutorials(bvids) {
  const [tutorials, setTutorials] = React.useState([]);
  const ref = React.useRef({});

  React.useEffect(() => {
    ref.current.tutorials = tutorials;
  }, [tutorials]);

  React.useEffect(() => {
    if (!bvids || !bvids.length) return;

    const token = Token.load();
    bvids.map(tutorial => {
      tutorial.link = `https://www.bilibili.com/video/${tutorial.id}`;

      axios.post(`/terraform/v1/mgmt/bilibili`, {
        ...token, bvid: tutorial.id,
      }).then(res => {
        const data = res.data.data;
        tutorial.title = data.title;
        tutorial.desc = data.desc;
        tutorial.view = parseInt(data.stat.view);
        tutorial.like = parseInt(data.stat.like);
        tutorial.share = parseInt(data.stat.share);
        // Order by view desc.
        setTutorials([...ref.current.tutorials, tutorial].sort((a, b) => b.view - a.view));
      });
      return null;
    });
  }, [bvids]);

  return tutorials;
}

export default function TutorialsButton({tutorials, prefixLine}) {
  const [show, setShow] = React.useState(false);

  return (
    <>
      <div role='button' style={{display: 'inline-block'}}>
        <Icon.PatchQuestion onClick={() => setShow(!show)} />
      </div>
      {show && prefixLine && <p></p>}
      {show && tutorials.map((tutorial, index) => (
        <div key={index}>
          <Toast show={show} onClose={() => setShow(false)}>
            <Toast.Header>
              <img src={logo} className="rounded me-2" width={56} alt=''/>
              <strong className="me-auto">bilibili</strong>
              <span title='播放次数'><Icon.Play /> {tutorial.view}</span> &nbsp;
              <span title='点赞次数'><Icon.HandThumbsUp /> {tutorial.like}</span> &nbsp;
              <span title='分享次数'><Icon.Share /> {tutorial.share}</span> &nbsp;
              <small>by {tutorial.author}</small>
            </Toast.Header>
            <Toast.Body>
              <a href={tutorial.link} target='_blank' rel='noreferrer'>
                {tutorial.title}
              </a>
            </Toast.Body>
          </Toast>
          <p></p>
        </div>
      ))}
    </>
  );
}

export {useTutorials, TutorialsButton};

