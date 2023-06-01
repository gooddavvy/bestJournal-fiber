# What is BestJournal, and why is it "Best"?

BestJournal is an **amazing** journal application built with GoFiber and **blazing fast** app with neat structure. Don't worry, it's the truth.

First of all, when I say neat structure, I mean neatly/nicely structured with multiple folders and files. This is a complex application and it took a lot of work to make, so be sure you like it and don't forget to **give this repository a STAR!**

# How can I try the site out?

All you need to do is to make sure you have GoLang installed on your machine, clone or fork this repository, open it in your favourite code editor (I recommend Visual Studio Code), then create a `.env` file. Add an environment variable with a key of `PORT` and a value of the port you'd like for your app, it looks something like this:

```env
PORT="8080"
```

Also, Add an environment variable with a key of `REGISTER_KEY` and a value of an empty string:

```env
REGISTER_KEY=""
```

After doing that, you can run the following command in the terminal:

```shell
$ go run .
    ┌───────────────────────────────────────────────────┐
    │                   Fiber v2.46.0                   │
    │               http://127.0.0.1:8080               │
    │       (bound on host 0.0.0.0 and port 8080)       │
    │                                                   │
    │ Handlers ............. 5  Processes ........... 1 │
    │ Prefork ....... Disabled  PID ............. 22036 │
    └───────────────────────────────────────────────────┘
```

You can now visit the app in your terminal and enjoy! If there are any issues, don't hesitate to let me know in the issues section!
