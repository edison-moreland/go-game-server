import requests
from tqdm import tqdm

URL = "http://localhost:8080"
NUM_REQUESTS = 10000


def benchmark_server() -> float:
    total_seconds = 0
    for _ in tqdm(range(NUM_REQUESTS)):
        # Update agent location
        res = requests.get(url=f"{URL}/agents")
        total_seconds += res.elapsed.microseconds

    return total_seconds/NUM_REQUESTS

def main():
    print("Setting up benchmark")

    # Start benchmarking
    print("Starting benchmark...")
    average_time = benchmark_server()
    print(f"Request millisecond average: {average_time/1000}")


if __name__ == '__main__':
    main()