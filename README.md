# Qtum Bridge

Welcome to the official repository of the Qtum Cross-Chain Bridge, a solution designed to seamlessly connect the Qtum and Ethereum blockchains. This project fosters interoperability between blockchain networks, enabling users to transfer assets securely and efficiently across different chains.

## Deployment

1. Build the Docker image:

```bash
docker build -t qtum-bridge .
```

2. Create a config file based on config [example](./config.yaml)

3. Run the Docker container:

```bash
docker run -d -p <your_local_port>:<port_from_config.yaml> -e KV_VIPER_FILE=/config.yaml -v <local_path_to_config.yaml>:/config.yaml qtum-bridge run service
```