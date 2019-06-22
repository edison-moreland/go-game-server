import random

from locust import HttpLocust, TaskSet, task


def random_location() -> dict:
    return {
        "X": random.randint(0, 100),
        "Y": random.randint(0, 100),
        "Z": random.randint(0, 100)
    }

class UserBehavior(TaskSet):
    user_id: str

    def on_start(self):
        """ on_start is called when a Locust start before any task is scheduled """
        # Create new agent
        res = self.client.post("/agent")
        self.user_id = res.json()["id"]

    def on_stop(self):
        """ on_stop is called when the TaskSet is stopping """
        self.client.delete(f"/agent/{self.user_id}")

    @task(2)
    def get_all(self):
        self.client.get("/agents")

    @task(1)
    def update_agent(self):
        self.client.put(f"/agent/{self.user_id}", json=random_location())

class WebsiteUser(HttpLocust):
    task_set = UserBehavior
    min_wait = 0
    max_wait = 0