# Core

Core module is responsible for general bridging logic:
- Track transactions
- Tracking flows that weren’t finished and allowing users to continue or cancel them
- Ingesting data about bridge events for faster access

Bridge Core DB will store data that it collects from smart contracts of different networks. All this data is also stored in blockchain, so could be restored easily.