wira-frontend
- use vue js to do
- need to run these prompt
    - `pnpm install`
    - `pnpm run dev`

wira-backend
- use golang to do backend 
- need to run 
    - `go mod tidy`
    - `go run .`

DB
- use progresql
- table I created by these dataset
    ```
    Account: {acc_id,username, email, encrypted_password, secretkey_2fa}
    Character: {char_id, name, acc_id, class_id,}
    Scores: {score_id, char_id, reward_score}
    Session: {session_id, session_metadata, expiry_datatime}
    ```
