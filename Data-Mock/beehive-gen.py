import requests
import random
import json
from datetime import datetime, timedelta

# === Configuration ===
GET_URL = "https://postgrest.kaminjitt.com/beehive?select=beehive_id,beehivestartdate"
POST_URL = "https://api.kaminjitt.com/api/harvestlog"
BEARER_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYXBpX3VzZXIifQ.Nx05HavMTlQfNZB3UVXe1KjAtMf_V3v2o0UT6uIA4y8"  # Replace with your actual token

HEADERS = {
    "Authorization": f"Bearer {BEARER_TOKEN}",
    "Content-Type": "application/json; type=array"
}

JSON_FILENAME = "harvestlog_records.json"

# === Functions ===

def fetch_beehives():
    response = requests.get(GET_URL, headers={"Authorization": f"Bearer {BEARER_TOKEN}"})
    response.raise_for_status()
    return response.json()

def random_date(start_str, end_dt):
    start_dt = datetime.fromisoformat(start_str)
    if start_dt > end_dt:
        return start_str
    delta = end_dt - start_dt
    rand_dt = start_dt + timedelta(seconds=random.randint(0, int(delta.total_seconds())))
    return rand_dt.strftime('%Y-%m-%d')

def generate_records(beehives, count=10):
    now = datetime.now()
    records = []

    for _ in range(count):
        hive = random.choice(beehives)
        beehive_id = hive["beehive_id"]
        start_date = hive["beehivestartdate"]

        record = {
            "beehive_id": beehive_id,
            "harvestdate": random_date(start_date, now),
            "production": round(random.uniform(0.5, 10.0), 2)
        }
        records.append(record)

    return records

def save_records_to_file(records, filename=JSON_FILENAME):
    with open(filename, 'w', encoding='utf-8') as f:
        json.dump(records, f, indent=2, ensure_ascii=False)
    print(f"✅ Saved {len(records)} records to {filename}")

def post_records(records):
    response = requests.post(POST_URL, headers=HEADERS, json=records)
    if response.status_code in (200, 201):
        print("✅ Successfully posted harvest records")
    else:
        print(f"❌ Failed to post records: {response.status_code} {response.text}")

# === Main Execution ===

def main():
    beehives = fetch_beehives()
    records = generate_records(beehives, count=1000000)  # adjust count as needed
    save_records_to_file(records)
    post_records(records)

if __name__ == "__main__":
    main()
