import random

import requests
from tqdm import tqdm

URL = "http://localhost:8080"
NUM_REQUESTS = 10000

def random_location() -> dict:
    return {
        "X": random.randint(0, 100),
        "Y": random.randint(0, 100),
        "Z": random.randint(0, 100)
    }

def benchmark_server(agent_id: str) -> float:
    total_seconds = 0
    for _ in tqdm(range(NUM_REQUESTS)):
        # Update agent location
        res = requests.put(url=f"{URL}/agent/{agent_id}", json=random_location())
        total_seconds += res.elapsed.microseconds

    return total_seconds/NUM_REQUESTS

def main():
    print("Setting up benchmark")

    # Create agent
    res = requests.post(url=f"{URL}/agent")
    agent_id = res.json()["id"]
    print(f"Agent ID: {agent_id}")

    # Start benchmarking
    print("Starting benchmark...")
    average_time = benchmark_server(agent_id)
    print(f"Request millisecond average: {average_time/1000}")

    # Cleanup
    requests.delete(url=f"{URL}/agent/{agent_id}")



if __name__ == '__main__':
    main()