# StatsBuddy-Go

StatsBuddy-Go is a Go-based application designed to manage and query cricket match data. It integrates with OpenAI and Gemini for natural language processing and MongoDB for data storage.

## Features

- Retrieve player information by ID or name pattern.
- Store and manage detailed cricket match information.
- Generate MongoDB queries using natural language prompts.
- Format responses using OpenAI and Gemini.

## Project Structure

- `api/`: Contains HTTP handlers for the API.
- `config/`: Configuration management.
- `llm/`: Integration with OpenAI and Gemini for natural language processing.
- `model/`: Data models for cricket match information.
- `players/`: Functions to retrieve player information.
- `service/`: Business logic for processing queries.

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/ganeshk312/statsBuddy-go.git
    cd statsBuddy-go
    ```

2. Set up environment variables:
    ```sh
    export OPENAI_API_KEY=your_openai_api_key
    export GEMINI_API_KEY=your_gemini_api_key
    export MONGODB_URI=your_mongodb_uri
    export MONGODB_DATABASE=your_database_name
    export LLM_TYPE=openai # or gemini
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

## API Endpoints

- `POST /query`: Process a natural language query and return the result.

## Example Request

```sh
curl -X POST http://localhost:8080/query -d '{"prompt": "Find all players with name pattern 'John'"}' -H "Content-Type: application/json"
```

## License

This project is licensed under the MIT License.
