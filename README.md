# goApi

This project implements a service for processing thousands of images collected from retail stores. The service receives jobs with image URLs and store IDs, calculates the perimeter of each image, and simulates GPU processing with a random sleep time. It provides an API for submitting jobs and checking the status of the jobs.


## Assumptions

- A job can consist of multiple images and may take a few minutes to an hour to process.
- The `store_id` and `image_url` are assumed to be valid, and image download failures will be recorded.
- If a job fails due to invalid `store_id` or image URL issues, a detailed error response will be provided.

## Installing and Setting Up the Project

### Prerequisites
- Go 1.22+ installed on your machine
- Docker
  
### Setup Instructions

1.Clone the repository:

git clone https://github.com/kirshiv/goApi.git

2. Install dependencies:

go mod tidy

3.Running the application (without Docker):

To run the application without Docker, execute the following command:

go run main.go

4.Running the application with Docker:

docker build -t retail-app .


Then, run the container:

docker run -d -p 8080:8080 --name retail-container retail-app

5.Testing the application:

The project provides API endpoints to submit jobs and check the status. I have tested it on `Postman`. I have written the steps to test it using postman:

- To submit a job (POST request):
  POST http://localhost:8080/api/submit/
  Request payload:
  json:
  {
     "count": 2,
     "visits": [
        {
           "store_id": "S00339218",
           "image_url": ["https://www.gstatic.com/webp/gallery/2.jpg", "https://www.gstatic.com/webp/gallery/3.jpg"],
           "visit_time": "2025-01-01T12:00:00Z"
        },
        {
           "store_id": "S01408764",
           "image_url": ["https://www.gstatic.com/webp/gallery/3.jpg"],
           "visit_time": "2025-01-01T13:00:00Z"
        }
     ]
  }
  ```

- To get job status (GET request):
  ```
  GET http://localhost:8080/api/status?jobid=123
  ```

6. Log output:

Logs will be printed to the console.

## Work Environment (`that i have used`)

- Computer/Operating System: Windows
- Text Editor/IDE: Visual Studio Code
- Libraries: 
- `net/http` for building the HTTP API
- `time` for simulating random delays
- `encoding/json` for working with JSON payloads
- `docker` for containerization

## Future Improvements

If given more time, the following improvements could be considered:
- Error Handling and Retries: Implement automatic retries for failed image downloads.
- Database Integration: Store job statuses and results in a persistent database instead of memory.
- Scalability: Enhance the service to process jobs in parallel, especially when handling thousands of images.
- Rate Limiting: Implement rate limiting to control the number of requests and prevent overloading.