# Syb Sequencer

A service that watches and processes events from the Sybil Resistance contract.

## Setup Instructions

### Environment Configuration

Create a `.env` file in the root directory with the following variables:

```env
# Database Configuration
DB_HOST="localhost"
DB_PORT="5432"
DB_USER=""
DB_PASSWORD=""
DB_NAME=""

# Ethereum Connection
ETHEREUM_RPC="wss://ethereum-rpc"
CONTRACT_ADDRESS="0x8A6d806C4d88c814Ad3d11125abcdb7410CCaa14"

# API Configuration
API_PORT="8080"
```

### Starting the Sequencer

1. Navigate to the sequencer directory:
   ```bash
   cd sequencer
   ```

2. Install dependencies and run:
   ```bash
   go mod tidy
   go run main.go
   ```

The sequencer will now listen to events emitted from the [contract on Sepolia](https://sepolia.etherscan.io/address/0x8A6d806C4d88c814Ad3d11125abcdb7410CCaa14).

## Testing with Event Triggers

### Triggering Contract Events

1. Checkout the feature branch in the MVP repository:
   ```bash
   git checkout feat/script-syb-sequencer
   ```

2. Navigate to the contracts directory:
   ```bash
   cd tokamak-sybil-resistance-mvp/contracts
   ```

3. Update your environment file with:
   ```env
   SYBIL_CONTRACT_ADDRESS="0x8A6d806C4d88c814Ad3d11125abcdb7410CCaa14"
   ```

4. Run the test script to emit events:
   ```bash
   forge script script/TestSyncEvents.s.sol:TestSyncEvents --rpc-url {YOUR_RPC_URL} --private-key {YOUR_PRIVATE_KEY} --broadcast
   ```

### Verifying Events

Query the API to check synced transactions:
```
GET {host}/api/v1/transactions
```

## Current Implementation

The sequencer currently operates as a real-time watcher that:
- Monitors events emitted from the smart contract
- Syncs transactions when events are detected
- Forges transactions after synchronization

> **Note:** The sequencer must be running when transactions are emitted to capture them successfully.

## Roadmap

We plan to enhance the sequencer to retroactively process events that occurred while the watcher was offline.