import {useNavigate, useSearchParams} from "react-router-dom";
import Container from "react-bootstrap/Container";
import React from "react";
import {Token, Errors} from "../utils";
import axios from "axios";
import {Row, Col, Card, Button, Form} from "react-bootstrap";
import UpgradeConfirmButton from '../components/UpgradeConfirmButton';
import * as semver from 'semver';
import * as moment from 'moment';

export default function System() {
  const navigate = useNavigate();
  const [status, setStatus] = React.useState();
  const [srs, setSRS] = React.useState();
  const [hooks, setHooks] = React.useState();
  const [tencent, setTencent] = React.useState();
  const [prometheus, setPrometheus] = React.useState();
  const [nodeExporter, setNodeExporter] = React.useState();
  const [strategyAutoUpgrade, setStrategyAutoUpgrade] = React.useState();
  const [userToggleStrategy, setUserToggleStrategy] = React.useState();
  const [searchParams] = useSearchParams();
  const [allowManuallyUpgrade, setAllowManuallyUpgrade] = React.useState();
  const [allowDisableContainer, setAllowDisableContainer] = React.useState();
  const [refreshContainers, setRefreshContainers] = React.useState();

  React.useEffect(() => {
    const allowManuallyUpgrade = searchParams.get('allow-manual') === 'true';
    console.log(`?allow-manual=true|false, current=${allowManuallyUpgrade}, Whether allow manually upgrade`);
    setAllowManuallyUpgrade(allowManuallyUpgrade);
  }, [searchParams]);

  React.useEffect(() => {
    const allowDisableContainer = searchParams.get('allow-disable') === 'true';
    console.log(`?allow-disable=true|false, current=${allowDisableContainer}, Whether allow disable container`);
    setAllowDisableContainer(allowDisableContainer);
  }, [searchParams]);

  const onStatus = (status) => {
    setStrategyAutoUpgrade(status.strategy === 'auto');
    setStatus(status);
  };

  const handleUpgradeStrategyChange = (e) => {
    if (strategyAutoUpgrade && !window.confirm(`关闭自动更新，将无法及时修复缺陷。\n是否确认关闭?`)) {
      e.preventDefault();
      return;
    }

    const token = Token.load();
    axios.post('/terraform/v1/mgmt/strategy', {
      ...token,
    }).then(res => {
      setUserToggleStrategy(!userToggleStrategy);
      console.log(`Strategy: Change ok`);
    }).catch(e => {
      const err = e.response.data;
      if (err.code === Errors.auth) {
        alert(`Token过期，请重新登录，${err.code}: ${err.data.message}`);
        navigate('/routers-logout');
      } else {
        alert(`服务器错误，${err.code}: ${err.data.message}`);
      }
    });
  };

  React.useEffect(() => {
    const token = Token.load();
    axios.post('/terraform/v1/mgmt/containers', {
      ...token, action: 'query',
    }).then(res => {
      const containers = res.data.data;
      containers.filter(container => {
        if (container.name === 'srs-server') setSRS(container);
        if (container.name === 'srs-hooks') setHooks(container);
        if (container.name === 'tencent-cloud') setTencent(container);
        if (container.name === 'prometheus') setPrometheus(container);
        if (container.name === 'node-exporter') setNodeExporter(container);
        return null;
      });
      console.log(`SRS: Query ok, containers are ${JSON.stringify(containers)}`);
    }).catch(e => {
      const err = e.response.data;
      if (err.code === Errors.auth) {
        alert(`Token过期，请重新登录，${err.code}: ${err.data.message}`);
        navigate('/routers-logout');
      } else {
        alert(`服务器错误，${err.code}: ${err.data.message}`);
      }
    });
  }, [navigate, refreshContainers]);

  const handleContainerChange = (container) => {
    const token = Token.load();
    axios.post('/terraform/v1/mgmt/containers', {
      ...token, action: 'enabled', name: container.name, enabled: !container.enabled,
    }).then(res => {
      console.log(`SRS: Update ok, enabled=${!container.enabled}`);
      setRefreshContainers(Math.random());
    }).catch(e => {
      const err = e.response.data;
      if (err.code === Errors.auth) {
        alert(`Token过期，请重新登录，${err.code}: ${err.data.message}`);
        navigate('/routers-logout');
      } else {
        alert(`服务器错误，${err.code}: ${err.data.message}`);
      }
    });
  };

  return (
    <>
      <Container>
        <Row>
          <Col xs lg={3}>
            <Card style={{ width: '18rem', marginTop: '16px' }}>
              <Card.Header>SRS服务器</Card.Header>
              <Card.Body>
                <Card.Text as={Col}>
                  容器名：{srs?.name} <br/>
                  容器ID：{srs?.container?.ID} <br/>
                  状态：{srs?.container.State} {srs?.container.Status}
                  <p></p>
                </Card.Text>
                <div style={{display: 'inline-block'}}>
                  <Button className='disabled'>
                    重启
                  </Button> &nbsp;
                  <Button className='disabled'>
                    升级
                  </Button> &nbsp;
                  <MgmtUpdateContainer
                    allow={allowDisableContainer}
                    enabled={srs?.enabled}
                    onClick={() => handleContainerChange(srs)}
                  />
                </div>
              </Card.Body>
            </Card>
          </Col>
          <Col xs lg={3}>
            <Card style={{ width: '18rem', marginTop: '16px' }}>
              <Card.Header>SRS回调</Card.Header>
              <Card.Body>
                <Card.Text as={Col}>
                  容器名：{hooks?.name} <br/>
                  容器ID：{hooks?.container?.ID} <br/>
                  状态：{hooks?.container.State} {hooks?.container.Status}
                  <p></p>
                </Card.Text>
                <div style={{display: 'inline-block'}}>
                  <Button className='disabled'>
                    重启
                  </Button> &nbsp;
                  <Button className='disabled'>
                    升级
                  </Button> &nbsp;
                  <MgmtUpdateContainer
                    allow={allowDisableContainer}
                    enabled={hooks?.enabled}
                    onClick={() => handleContainerChange(hooks)}
                  />
                </div>
              </Card.Body>
            </Card>
          </Col>
          <Col xs lg={3}>
            <Card style={{ width: '18rem', marginTop: '16px' }}>
              <Card.Header>腾讯云</Card.Header>
              <Card.Body>
                <Card.Text as={Col}>
                  容器名：{tencent?.name} <br/>
                  容器ID：{tencent?.container?.ID} <br/>
                  状态：{tencent?.container.State} {tencent?.container.Status}
                  <p></p>
                </Card.Text>
                <div style={{display: 'inline-block'}}>
                  <Button className='disabled'>
                    重启
                  </Button> &nbsp;
                  <Button className='disabled'>
                    升级
                  </Button> &nbsp;
                  <MgmtUpdateContainer
                    allow={allowDisableContainer}
                    enabled={tencent?.enabled}
                    onClick={() => handleContainerChange(tencent)}
                  />
                </div>
              </Card.Body>
            </Card>
          </Col>
          <Col xs lg={3}>
            <Card style={{ width: '18rem', marginTop: '16px' }}>
              <Card.Header>Prometheus监控</Card.Header>
              <Card.Body>
                <Card.Text as={Col}>
                  容器名：{prometheus?.name} <br/>
                  容器ID：{prometheus?.container?.ID} <br/>
                  状态：{prometheus?.container.State} {prometheus?.container.Status}
                  <p></p>
                </Card.Text>
                <div style={{display: 'inline-block'}}>
                  <Button className='disabled'>
                    重启
                  </Button> &nbsp;
                  <Button className='disabled'>
                    升级
                  </Button> &nbsp;
                  <MgmtUpdateContainer
                    allow={allowDisableContainer}
                    enabled={prometheus?.enabled}
                    onClick={() => handleContainerChange(prometheus)}
                  />
                </div>
              </Card.Body>
            </Card>
          </Col>
          <Col xs lg={3}>
            <Card style={{ width: '18rem', marginTop: '16px' }}>
              <Card.Header>NodeExporter(节点监控)</Card.Header>
              <Card.Body>
                <Card.Text as={Col}>
                  容器名：{nodeExporter?.name} <br/>
                  容器ID：{nodeExporter?.container?.ID} <br/>
                  状态：{nodeExporter?.container.State} {nodeExporter?.container.Status}
                  <p></p>
                </Card.Text>
                <div style={{display: 'inline-block'}}>
                  <Button className='disabled'>
                    重启
                  </Button> &nbsp;
                  <Button className='disabled'>
                    升级
                  </Button> &nbsp;
                  <MgmtUpdateContainer
                    allow={allowDisableContainer}
                    enabled={nodeExporter?.enabled}
                    onClick={() => handleContainerChange(nodeExporter)}
                  />
                </div>
              </Card.Body>
            </Card>
          </Col>
          <Col xs lg={3}>
            <Card style={{ width: '18rem', marginTop: '16px' }}>
              <Card.Header>管理后台</Card.Header>
              <Card.Body>
                <Card.Text as={Col}>
                  你的版本: {status?.version} <br/>
                  稳定版本: {status?.releases?.stable} &nbsp;
                  <Form.Check
                    type='switch'
                    label='自动更新'
                    style={{display: 'inline-block'}}
                    title='是否自动更新到稳定版本'
                    disabled={!allowManuallyUpgrade}
                    defaultChecked={strategyAutoUpgrade}
                    onClick={(e) => handleUpgradeStrategyChange(e)}
                  />
                  <br/>
                  最新版本: <a href='https://github.com/ossrs/srs/issues/2856#changelog' target='_blank' rel='noreferrer'>{status?.releases?.latest}</a>
                  <p></p>
                </Card.Text>
                <MgmtUpgradeButton onStatus={onStatus}/>
              </Card.Body>
            </Card>
          </Col>
        </Row>
      </Container>
    </>
  );
}

function MgmtUpdateContainer({allow, enabled, onClick}) {
  const handleClick = (e) => {
    if (enabled && !window.confirm(`禁用容器，将会停止服务。\n服务将不可用，且重启服务器也不会启动。\n是否确认禁用?`)) {
      e.preventDefault();
      return;
    }
    onClick();
  };

  return (
    <Button
      className={allow ? '' : 'disabled'}
      variant={enabled ? 'danger' : 'success'}
      onClick={(e) => handleClick(e)}
    >
      {enabled ? '禁用' : '启用'}
    </Button>
  );
}

function MgmtUpgradeButton({onStatus}) {
  const [startingUpgrade, setStartingUpgrade] = React.useState();
  const [isUpgrading, setIsUpgrading] = React.useState();
  const [releaseAvailable, setReleaseAvailable] = React.useState();
  const [upgradeDone, setUpgradeDone] = React.useState();
  const [progress, setProgress] = React.useState(120);

  // For callback to use state.
  const ref = React.useRef({});
  React.useEffect(() => {
    ref.current.startingUpgrade = startingUpgrade;
    ref.current.progress = progress;
    ref.current.upgradeDone = upgradeDone;
  }, [startingUpgrade, progress, upgradeDone]);

  React.useEffect(() => {
    const refreshMgmtStatus = () => {
      const token = Token.load();
      axios.post('/terraform/v1/mgmt/status', {
        ...token,
      }).then(res => {
        const status = res.data.data;

        // Normally state.
        setIsUpgrading(status.upgrading);
        onStatus(status);

        // Whether upgrade is available.
        if (status && status.releases && status.releases.latest) {
          setReleaseAvailable(semver.lt(status.version, status.releases.latest));
        }

        // If upgradeDone is false, we're in the upgrading progress, so it's done when upgrading changed to false.
        if (ref.current.upgradeDone === false && !status.upgrading && ref.current.progress < 120) {
          setStartingUpgrade(false);
          setUpgradeDone(true);
        }

        // If state not set, but already upgrading, it's restore from the previous state.
        if (status.upgrading && ref.current.upgradeDone === undefined && ref.current.startingUpgrade === undefined) {
          setStartingUpgrade(true);
          setUpgradeDone(false);
        }

        console.log(`${moment().format()}: Status: Query ok, startingUpgrade=${ref.current.startingUpgrade}, upgradeDone=${ref.current.upgradeDone}, status=${JSON.stringify(status)}`);
      }).catch(e => {
        console.log('ignore any error during status', e);
      });
    };

    refreshMgmtStatus();
    const timer = setInterval(() => refreshMgmtStatus(), 5000);
    return () => clearInterval(timer);
  }, [startingUpgrade]);

  const handleStartUpgrade = () => {
    if (isUpgrading) return;

    setUpgradeDone(false);
    setStartingUpgrade(true);
    setProgress(120);

    const token = Token.load();
    axios.post('/terraform/v1/mgmt/upgrade', {
      ...token,
    }).then(res => {
      console.log(`upgrade ok, ${JSON.stringify(res.data.data)}`);
    }).catch(e => {
      console.log('ignore any error during upgrade', e);
    });
  };

  React.useEffect(() => {
    if (!isUpgrading) return;
    const timer = setInterval(() => {
      if (ref.current.progress <= 0) return;
      if (ref.current.progress <= 10) setUpgradeDone(true);
      setProgress(ref.current.progress - 1);
    }, 1000);
    return () => clearInterval(timer);
  }, [isUpgrading]);

  React.useEffect(() => {
    if (!upgradeDone) return;
    alert('升级成功，请刷新页面');
  }, [upgradeDone]);

  return (
    <UpgradeConfirmButton
      releaseAvailable={releaseAvailable}
      upgrading={startingUpgrade || isUpgrading}
      progress={`${progress}s`}
      onClick={handleStartUpgrade}
      text='升级'
    >
      <p>
        升级管理后台，并且可能造成
        <span className='text-danger'><strong>系统不可用</strong></span>，
        确认继续升级么？
      </p>
    </UpgradeConfirmButton>
  );
}

