# Booking and Reservations

- Built in Go version 1.15
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)


## Star db

```
docker run --name bookingsgo -e POSTGRES_PASSWORD=mysecret -e POSTGRES_DB=bookings -p 5432:5432 -d postgres
docker container start bookingsgo
```

## Credits and Inspiration:
https://www.udemy.com/course/building-modern-web-applications-with-go/

Initial commit started from: https://github.com/gandra/helloworld-web-go-trevor-sawler
