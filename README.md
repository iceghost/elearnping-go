# E-learnping! Server

A web server for querying updates on BKeL.

```
$ curl \
    --location --request GET 'https://elearnping-go.fly.dev/api/sites' \
    --header 'Authorization: Bearer abc123xyz456' \

[
    {
        "site": {
            "id": 62848,
            "fullname": "C\u1ea5u tr\u00fac r\u1eddi r\u1ea1c cho khoa h\u1ecdc m\u00e1y t\u00ednh (CO1007)_Video",
            "groupid": 173774
        },
        "from": "2022-05-07T19:00:00+07:00",
        "to": "2022-05-15T23:36:01.045185994+07:00",
        "updates": [
            {
                "module": {
                    "id": 803748,
                    "name": "Conditional Prob 1",
                    "modname": "quiz"
                },
...
```

## Instruction

1. Install Redis for cache database
2. Make a `.env` file with Redis connection details (see `.env.example`)
3. Run Redis server

```bash
$ redis-server
```

4. Run server

```bash
$ go run .
```

5. Make GET requests to server, with your moodle token as Bearer token on Authorization header.

```bash
$ curl \
    --location --request GET 'localhost:8080/api/sites' \
    --header 'Authorization: Bearer {{YOUR_MOODLE_TOKEN_HERE}}' \
    | python3 -m json.tool
```

## Where to find your Moodle token?

Login into e-learning site. Inside the menu on the arrow left of your avatar,
click "Tùy chọn"

![](images/step2.jpg)

Click "Security keys".

![](images/step3.jpg)

Your token is the "Moodle mobile web service" one.

![](images/step4.jpg)
