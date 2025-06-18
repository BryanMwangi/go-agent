# DevLogs

## 2025-06-18, 18:30

We start our journey with the `go-agent` project. The task is to create a CLI tool that can be used to interact with an LLM, similar to `claude-code`. The basis of this project is the `claude-code` project, which is a CLI tool that allows users to interact with a language model using a command-line interface.

The key challenge is `claude-code` is written in JavaScript and we ought to write a Go version more compatible with other LLM providers other than Anthropic. The main goal is to create a CLI tool that can be used to interact with any LLM provider.

For the sake of simplicity, we will mirror some of the functionalities of `claude-code` in the `go-agent` project but use OpenAI as the LLM provider to spice it up a bit.

### Goals

- The agent should have access to tools, ideally reading from the file system, searching the internet, access to grep and other terminal commands, and running code.
- Perhaps being able to talk to a language server would help as well for coding tasks?
- An MCP client, that can call into MCP servers that it’s linked to would allow users to extend the tools and resources available to the agent.
- The agent should be able to access the file system safely and execute terminal commands safely; ideally the agent should be limited to working directory and have safeguards to prevent files from being destroyed or leaked; perhaps landlock the process for OS’s that support it, or other layers like regex over commands / asking for permissions from the user?
- The agent should be able to run bash and python scripts it generates in a safe way; perhaps in a sandbox, a WASM module, a docker container, a VM, etc.
- Be able to use an LLM other than clause code; support other APIs and models. We probably want model-specific system prompts; different models may need to be prompted and tweaked in different ways to get the most out of them.
- Some sort of context-compaction would be cool, in order to keep context smaller without losing important information.
- Some sort of long term memory, like writing a to do list or documentation, would be nice as well
- You could use a Golang terminal UI, TUI, to create a more interesting and dynamic interface than the basic CLI Claude has.

### Realistic Targets

The above goals were set but due to a time limitation of 10 hours, we have to narrow it down to the following:

- The agent should be able to read and write files to the file system
- The agent should be able to execute terminal commands
- The agent should be able to run bash scripts

## In the beginning

### 18:45

We start by creating a new Go module called `go-agent` and setting up the basic structure of the project. Using `claude-code` as our reference, we know that there are 3 main components our application needs to get started:

1. The Terminal UI
2. The Agent or client that will interact with the various APIs
3. Basic tools to allow simple prompts and responses before elevating the scope to more complex tasks

We will start by creating a `terminal` package that will contain the Terminal UI. We will use `tview` as our terminal UI library. [`tview`](https://github.com/rivo/tview) is a Go library for building terminal user interfaces.

### 19:00

Next, we will analyze the deobfuscated version of `claude-code` and identify the main components of the application. The main goal is to find how natural language is processed in order to know what actions the AI agent can execute in the code base.

We were able to find one particular function that is called for processing the user's input that being `processUserInput`. This function is responsible for handling the user's input to identify if the input calls for a '/' command. What we do not know is what happens in the case a '/' command is not recognized.

In order to solve this problem, we will first focus the flow of the application keeping in mind that the best way to do this is by using known commands. Later, when no command is defined, we will use the LLM to first analyze the user's input, dynamically generate the required commands, and then execute them.

What this allows us to do is using the LLM to determine the user's intent without the need for the user to explicitly define the command.

## 19:27 - 20:30

Reading OpenAI's documentation with curl implementations as we will be using the raw API.

## 20:30

We now start writing the code from the config file alongside the llm and terminal packages. Our implementation will use API Keys to avoid oauth authentication. This speeds up our workflow.

## 20:49-21:11

Switched terminal rendering to use `promptui` instead of `tview`. This is because I perceieved `tview` as a bit too over-engineered for this project. I also wanted to make the experience more similar to `claude-code`.

## 21:11-21:30

Small break to think about authentication flow and later command flow and prompts.

## 21:30-23:00

Set up authentication flow and prompts. Here we establish a session and create or load the config file. We also verified that the API key is valid by calling the OpenAI API. We stored the user's session including the user's name and model they selected. When the app is restarted, we can continue from previous working directory and session or
we ask the user to select a new working directory. It is important to note that for the sake of simplicity at this time, a new session will force the user to reauthenticate.

## 23:10-[to be determined]

Small break to think about command flow and prompts.
