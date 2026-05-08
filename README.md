# LearnGoNotificationApp

A Go project I'm building as I learn the language — applied to a
real-world notification system architecture.

Rather than following a tutorial, I'm using Go concepts as I learn them
to build something that reflects real-world patterns. Each new concept
gets applied directly to this project rather than living in a throwaway
exercise file.

## What it does

Models a customer outreach notification system where recipients receive
reminders or blast notifications through different channels (SMS, email).
The system tracks outreach status per recipient and handles failures
gracefully.

## Architecture

```
LearnGoNotificationApp/
    main.go              — orchestration, main loop, result categorization
    go.mod
    enums/
        outreachStatus.go  — typed OutreachStatus enum (NotSent, Failed, Success)
        outreachType.go    — typed OutreachType enum (Reminder, Blast)
    models/
        notificationSender.go — NotificationSender interface
        person.go          — Person struct with status tracking
    senders/
        emailSender.go     — Email implementation
        smsSender.go       — SMS implementation
    utils/
        maputils.go        — generic FilterMap utility for map filtering
    workers/
        notificationWorker.go — worker pool, processes notifications from channel
```

## Concepts applied so far

- Package organization — models, senders, enums as separate packages
- Interface-driven architecture — pluggable sender implementations
- Typed enums — OutreachType and OutreachStatus prevent invalid states at compile time
- Error handling — multiple return values, no try/catch
- Structs and methods — Person with constructor and Stringer
- Exported vs unexported identifiers — Go's access control model
- Maps, slices, for range loops — core data structures in practice
- Pointers — why the map stores *Person not Person
- Goroutines + channels + WaitGroups — worker pool pattern for concurrent processing

## Learning log

| Concept | Applied as |
|---|---|
| Variables, types, control flow | Notification routing and status checks |
| Functions + multiple return values | sendReminder returning (bool, error) |
| Structs + methods | Person model with String() formatter |
| Interfaces | NotificationSender with SMS and Email implementations |
| Package organization | models / senders / enums separation |
| Typed enums | OutreachType replacing raw strings |
| Goroutines + WaitGroups | Coming next — concurrent notification processing |
| Generics + higher-order functions | FilterMap utility for filtering notification results |
| Goroutines + channels + WaitGroups | Worker pool for concurrent notification processing |

## What's coming

- Tests — table-driven tests for FilterMap, processNotification, and worker behavior
- HTTP integration for real SMS/email providers
- Config file for recipient data instead of hardcoded values
- Additional outreach channels (push notifications)

## Why this project

Notification systems are a common real-world pattern — queuing, routing,
status tracking, failure handling, pluggable integrations. Building
around a familiar architecture means I can focus on learning Go without
also figuring out the problem domain from scratch.

## Stack

- Go 1.21+
- No external dependencies — standard library only