# Milestones

List of achievements, and features still pending to implement.

## TODO

- Add character creation.
- Add basic gameplay teleport
- Implement all packets
  - 0x11 Character status
  - 0x17 Update health bar
  - 0x20 Update player
  - 0x4F Light level
  - 0x55 Login complete
  - 0x5B Implement server time
  - 0x6D Play music
  - 0x72 War mode
  - 0x78 Update object
  - 0xBF Implement all subcommands
    - 0x08 Map change
    - 0x18 Apply world patches
    - 0x19 Extended stats

## Done

- Mono port: Server only runs on one port, and can infer what encryption to use.
- Gameplay cryptography: Both decrypt and compress+encrypt works.