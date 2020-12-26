import pytest
import testinfra
import requests
import os

api_address = 'http://127.0.0.1:%s' % os.environ['PORT']
systemd_host = testinfra.get_host('docker://%s' % os.environ['CONTAINER_ID'])

def test_init():
    systemd_host.run('systemctl start testing-unit.service')
    status = systemd_host.run('systemctl -q is-active testing-unit.service')
    assert status.rc == 0

def test_get_existing_units():
    req = requests.get('%s/api/v1/units' % api_address)
    assert len(req.json()) >= 83

def test_get_existing_unit_response():
    req = requests.get('%s/api/v1/units/testing-unit.service' % api_address)
    expected_keys = ['Name', 'Description', 'LoadState', 'ActiveState', 'SubState',
                     'Followed', 'Path', 'JobId', 'JobType', 'JobPath']
    result = req.json()
    assert len(result) == 10
    for key in expected_keys:
        assert key in result

def test_get_existing_unit_status():
    req = requests.get('%s/api/v1/units/testing-unit.service' % api_address)
    assert req.json()['ActiveState'] == 'active'

def test_post_existing_unit_stop():
    requests.post('%s/api/v1/units/testing-unit.service/stop' % api_address)

    status = systemd_host.run('systemctl -q is-active testing-unit.service')
    assert status.rc != 0

    req = requests.get('%s/api/v1/units/testing-unit.service' % api_address)
    assert req.json()['ActiveState'] != 'active'

def test_post_existing_unit_start():
    requests.post('%s/api/v1/units/testing-unit.service/start' % api_address)

    status = systemd_host.run('systemctl -q is-active testing-unit.service')
    assert status.rc == 0

    req = requests.get('%s/api/v1/units/testing-unit.service' % api_address)
    assert req.json()['ActiveState'] == 'active'

def test_get_nonexisting_unit_status():
    req = requests.get('%s/api/v1/units/non-exist.service' % api_address)
    assert req.status_code == 404

def test_post_nonexisting_unit_stop():
    req = requests.post('%s/api/v1/units/non-exist.service/stop' % api_address)
    assert req.status_code == 404

def test_post_nonexisting_unit_start():
    req = requests.post('%s/api/v1/units/non-exist.service/start' % api_address)
    assert req.status_code == 404
