import http.server
import json
import click
import time

from .problem_info import ParsedProblem

class ProblemHandler(http.server.BaseHTTPRequestHandler):
    server: "TimedHTTPServer"

    def do_POST(self):
        content_length = int(self.headers['Content-Length'])
        body = self.rfile.read(content_length)
        
        problem_data = json.loads(body)
        parsed_problem = ParsedProblem(**problem_data)
        click.echo(f"Received problem: {parsed_problem.name}")

        self.server.batch.append(parsed_problem)
        self.server.last_problem_time = time.time()

        self.send_response(200)
        self.end_headers()

    def log_message(self, *args):
        pass

class TimedHTTPServer(http.server.HTTPServer):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.batch: list[ParsedProblem] = []
        self.last_problem_time = None
        self.timeout = 1

    def serve_with_timeout(self, no_problem_timeout: int = 5) -> list[ParsedProblem]:
        start_time = time.time()
        first_problem_wait = 180

        while True:
            try:
                if not len(self.batch):
                    if time.time() - start_time > first_problem_wait:
                        click.echo("\nNo problems received within 3 minutes. Shutting down...")
                        break
                    self.handle_request()
                    continue

                if time.time() - self.last_problem_time > no_problem_timeout:
                    click.echo("\nNo new problems received for 5 second. Processing complete.")
                    break

                self.handle_request()
            except TimeoutError:
                continue

        return self.batch
