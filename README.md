## TwinRoom API client

This repository hosts a sample client application for the TwinRoom API, created by [ykimura517](https://twitter.com/yk_llm_gpt).   
It demonstrates how to connect to the TwinRoom API endpoint using Server-Sent Events (SSE).   
By default, the response texts from the API are concatenated into a single string.   
However, for actual implementations, you have the option to process the responses either in the same manner or as a stream.

## Usage

### Register to TwinRoom 

To sign up for TwinRoom, please get in touch with us.

https://go-spiral.ai/contact/  

Once your organization account is set up, team members will gain access to the dashboard.

### Create API Key
Within the TwinRoom dashboard, you can create your AI character and generate its API key.   
Configure the TwinRoom API endpoint and your API key in the .env file (refer to .env.sample for an example).

### Build Docker Container
Execute the following commands:
```
docker build -t twinroom-sample .
```
```
docker run -p 8083:8083 --env-file ./.env twinroom-sample
```

### Check AI Response

After the Docker container is up and running, navigate to http://localhost:8083/stream in your browser. The page will display response texts, and if voice functionality is enabled, the audio files will be stored in the voice folder.

That's all!
Happy hacking! :)