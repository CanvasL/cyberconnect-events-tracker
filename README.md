# CyberConnect Events Tracker

This project is designed to solve two information retrieval problems with the [CyberConnect](https://github.com/cyberconnecthq/cybercontracts) smart contract design. Firstly, it is difficult to obtain all `ProfileIds` owned by a specific wallet address. Secondly, it is challenging to retrieve the `collect information` set by the publisher when registering an essence on the middleware. Additionally, the [CyberConnect API](https://cyberconnect-docs-v2.vercel.app/) design is not very user-friendly, especially when integrating it into an SDK. It requires prior login verification, which is not ideal.

To address these issues, we have set up a backend that listens to contract events on both the `BSC mainnet` and `BSC Testnet`. The events monitored are `CollectPaidMwSet` and `CreateProfile` , and the information is stored in a MySQL database. The backend exposes an API that can be used by external clients to retrieve the stored information.

## Installation

To install and run the project, follow these steps:

1. Clone the repository
2. Install the required dependencies using go get
3. Set up a MySQL database and create the tables using the provided SQL script
4. Set the environment variables for the database connection and the BSC node URL
5. Run the project using go run main.go

## API

The following endpoints are exposed by the backend:

- GET `/api/v1/{chainId}/profiles-info/{address}` : Retrieves all ProfileIds owned by the specified wallet address.
- GET `/api/v1/{chainId}/collect-info/{profileId}/{essenceId}` : Retrieves the collect information set by the publisher when registering an essence on the middleware.

## Conclusion

By setting up this backend, we have made it easier for developers to interact with the CyberConnect smart contract and retrieve the necessary information. If you have any questions or suggestions, please feel free to open an issue or submit a pull request.
