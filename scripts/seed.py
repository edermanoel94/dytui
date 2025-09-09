import os
import time
import random
import string
from uuid import uuid4
from datetime import datetime, timezone

import boto3
from botocore.exceptions import ClientError

ENDPOINT = os.getenv("DYNAMO_ENDPOINT", "http://localhost:4566")
REGION = os.getenv("AWS_DEFAULT_REGION", "us-east-1")
NUM_ITEMS = int(os.getenv("NUM_ITEMS", "20"))

dynamo = boto3.client("dynamodb", region_name=REGION, endpoint_url=ENDPOINT)

TABLES = [f"demo_table_{i}" for i in range(1, 6)]

def ensure_table(name: str):
    try:
        dynamo.describe_table(TableName=name)
        print(f"[=] Table exists: {name}")
        return
    except ClientError as e:
        if e.response["Error"]["Code"] != "ResourceNotFoundException":
            raise

    print(f"[+] Creating table: {name}")
    dynamo.create_table(
        TableName=name,
        AttributeDefinitions=[
            {"AttributeName": "id", "AttributeType": "S"},
        ],
        KeySchema=[{"AttributeName": "id", "KeyType": "HASH"}],
        BillingMode="PAY_PER_REQUEST",
        Tags=[{"Key": "env", "Value": "local"}, {"Key": "owner", "Value": "seed"}],
    )

    waiter = dynamo.get_waiter("table_exists")
    waiter.wait(TableName=name)
    print(f"[✓] Table ready: {name}")

def rand_str(n=10):
    return "".join(random.choices(string.ascii_letters + string.digits, k=n))

def rand_tags():
    base = ["alpha", "beta", "gamma", "delta", "omega", "kappa", "zeta", "tau", "phi", "psi"]
    k = random.randint(1, 4)
    return random.sample(base, k=k)

def put_items(table: str, n: int):
    print(f"[>] Inserting {n} items into {table} ...")
    for _ in range(n):
        item = {
            "id": {"S": str(uuid4())},
            "name": {"S": rand_str(12)},
            "score": {"N": str(random.randint(0, 1000))},
            "active": {"BOOL": random.choice([True, False])},
            "created_at": {"S": datetime.now(timezone.utc).isoformat()},
            "tags": {"L": [{"S": t} for t in rand_tags()]},
            "meta": {"M": {
                "source": {"S": random.choice(["seed.py", "fixture", "import"])},
                "version": {"N": str(random.randint(1, 5))}
            }}
        }
        dynamo.put_item(TableName=table, Item=item)
    print(f"[✓] Done: {table}")

def main():
    # simple backoff in case compose healthcheck passes just before full readiness
    for i in range(10):
        try:
            health = dynamo.list_tables()
            if "TableNames" in health:
                break
        except Exception:
            time.sleep(1 + i * 0.2)

    for t in TABLES:
        ensure_table(t)
        put_items(t, NUM_ITEMS)

    print("[✅] Seeding complete.")

if __name__ == "__main__":
    main()
