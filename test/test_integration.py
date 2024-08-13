import json

import pytest
import requests


@pytest.fixture(scope="module")
def base_url():
    return "http://127.0.0.1:8080"

def test_simple_booking(base_url):
    headers = {
        "Content-Type": "application/json"
    }

    url_ping = f"{base_url}/ping"

    response = requests.get(url_ping, headers=headers)
    assert response.status_code == 200

    url_add_hotel = f"{base_url}/add_hotel"
    payload = {
        "hotel_id": "reddison",
        "rooms": ["lux"],
    }

    response = requests.post(url_add_hotel, data=json.dumps(payload), headers=headers)
    assert response.status_code == 201

    url_booking = f"{base_url}/orders"
    payload = {
        "hotel_id": "reddison",
        "room_id": "lux",
        "email": "guest@mail.ru",
        "from": "2030-01-02T00:00:00Z",
        "to": "2030-01-04T00:00:00Z"
    }

    response = requests.post(url_booking, data=json.dumps(payload), headers=headers)
    assert response.status_code == 201

    response = requests.post(url_booking, data=json.dumps(payload), headers=headers)
    assert response.status_code == 500