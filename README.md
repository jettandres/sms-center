# SMS Center

This server will store my personal sms messages in a lightweight sqlite database

# Potential usage:
- Store incoming SMS messages via `POST` endpoint
- Retrieve sms messages via `GET` endpoint
- Retrieve real-time database copy using `syncthing` installed in my linode instance

# Components:

- Main Phone
  - will delegate incoming SMS messages using Tasker and `curl` command
- Backup Phone
  - will have a dedicated RN app to view messages from the server
