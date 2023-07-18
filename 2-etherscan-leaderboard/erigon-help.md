NAME:
   erigon - erigon

USAGE:
   erigon [command] [flags]

VERSION:
   2.48.1-stable-674b77f0

COMMANDS:
   init       Bootstrap and initialize a new genesis block
   import     Import a blockchain file
   snapshots  Managing snapshots (historical data partitions)
   support    Connect Erigon instance to a diagnostics system for support
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --datadir value                                          Data directory for the databases (default: /Users/syjn99/Library/Erigon)
   --ethash.dagdir value                                    Directory to store the ethash mining DAGs (default: /Users/syjn99/Library/erigon-ethash)
   --snapshots                                              Default: use snapshots "true" for Mainnet, Goerli, Gnosis Chain and Chiado. use snapshots "false" in all other cases (default: true)
   --internalcl                                             enables internal consensus (default: false)
   --txpool.disable                                         experimental external pool and block producer, see ./cmd/txpool/readme.md for more info. Disabling internal txpool and block producer. (default: false)
   --txpool.locals value                                    Comma separated accounts to treat as locals (no flush, priority inclusion)
   --txpool.nolocals                                        Disables price exemptions for locally submitted transactions (default: false)
   --txpool.pricelimit value                                Minimum gas price (fee cap) limit to enforce for acceptance into the pool (default: 1)
   --txpool.pricebump value                                 Price bump percentage to replace an already existing transaction (default: 10)
   --txpool.accountslots value                              Minimum number of executable transaction slots guaranteed per account (default: 16)
   --txpool.globalslots value                               Maximum number of executable transaction slots for all accounts (default: 10000)
   --txpool.globalbasefeeslots value                        Maximum number of non-executable transactions where only not enough baseFee (default: 30000)
   --txpool.accountqueue value                              Maximum number of non-executable transaction slots permitted per account (default: 64)
   --txpool.globalqueue value                               Maximum number of non-executable transaction slots for all accounts (default: 30000)
   --txpool.lifetime value                                  Maximum amount of time non-executable transaction are queued (default: 3h0m0s)
   --txpool.trace.senders value                             Comma separared list of addresses, whoes transactions will traced in transaction pool with debug printing
   --txpool.commit.every value                              How often transactions should be committed to the storage (default: 15s)
   --prune value                                            Choose which ancient data delete from DB:
                                                            h - prune history (ChangeSets, HistoryIndices - used by historical state access, like eth_getStorageAt, eth_getBalanceAt, debug_traceTransaction, trace_block, trace_transaction, etc.)
                                                            r - prune receipts (Receipts, Logs, LogTopicIndex, LogAddressIndex - used by eth_getLogs and similar RPC methods)
                                                            t - prune transaction by it's hash index
                                                            c - prune call traces (used by trace_filter method)
                                                            Does delete data older than 90K blocks, --prune=h is shortcut for: --prune.h.older=90000.
                                                            Similarly, --prune=t is shortcut for: --prune.t.older=90000 and --prune=c is shortcut for: --prune.c.older=90000.
                                                            However, --prune=r means to prune receipts before the Beacon Chain genesis (Consensus Layer might need receipts after that).
                                                            If an item is NOT on the list - means NO pruning for this data.
                                                            Example: --prune=htc (default: "disabled")
   --prune.h.older value                                    Prune data older than this number of blocks from the tip of the chain (if --prune flag has 'h', then default is 90K) (default: 0)
   --prune.r.older value                                    Prune data older than this number of blocks from the tip of the chain (default: 0)
   --prune.t.older value                                    Prune data older than this number of blocks from the tip of the chain (if --prune flag has 't', then default is 90K) (default: 0)
   --prune.c.older value                                    Prune data older than this number of blocks from the tip of the chain (if --prune flag has 'c', then default is 90K) (default: 0)
   --prune.h.before value                                   Prune data before this block (default: 0)
   --prune.r.before value                                   Prune data before this block (default: 0)
   --prune.t.before value                                   Prune data before this block (default: 0)
   --prune.c.before value                                   Prune data before this block (default: 0)
   --batchSize value                                        Batch size for the execution stage (default: "256M")
   --bodies.cache value                                     Limit on the cache for block bodies (default: "268435456")
   --database.verbosity value                               Enabling internal db logs. Very high verbosity levels may require recompile db. Default: 2, means warning. (default: 2)
   --private.api.addr value                                 Erigon's components (txpool, rpcdaemon, sentry, downloader, ...) can be deployed as independent Processes on same/another server. Then components will connect to erigon by this internal grpc API. example: 127.0.0.1:9090, empty string means not to start the listener. do not expose to public network. serves remote database interface (default: "127.0.0.1:9090")
   --private.api.ratelimit value                            Amount of requests server handle simultaneously - requests over this limit will wait. Increase it - if clients see 'request timeout' while server load is low - it means your 'hot data' is small or have much RAM.  (default: 31872)
   --etl.bufferSize value                                   Buffer size for ETL operations. (default: "256MB")
   --tls                                                    Enable TLS handshake (default: false)
   --tls.cert value                                         Specify certificate
   --tls.key value                                          Specify key file
   --tls.cacert value                                       Specify certificate authority
   --state.stream.disable                                   Disable streaming of state changes from core to RPC daemon (default: false)
   --sync.loop.throttle value                               Sets the minimum time between sync loop starts (e.g. 1h30m, default is none)
   --bad.block value                                        Marks block with given hex string as bad and forces initial reorg before normal staged sync
   --http                                                   HTTP-RPC server (enabled by default). Use --http=false to disable it (default: true)
   --graphql                                                Enable the graphql endpoint (default: false)
   --http.addr value                                        HTTP-RPC server listening interface (default: "localhost")
   --http.port value                                        HTTP-RPC server listening port (default: 8545)
   --authrpc.addr value                                     HTTP-RPC server listening interface for the Engine API (default: "localhost")
   --authrpc.port value                                     HTTP-RPC server listening port for the Engine API (default: 8551)
   --authrpc.jwtsecret value                                Path to the token that ensures safe connection between CL and EL
   --http.compression                                       Enable compression over HTTP-RPC (default: false)
   --http.corsdomain value                                  Comma separated list of domains from which to accept cross origin requests (browser enforced)
   --http.vhosts value                                      Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts 'any' or '*' as wildcard. (default: "localhost")
   --authrpc.vhosts value                                   Comma separated list of virtual hostnames from which to accept Engine API requests (server enforced). Accepts 'any' or '*' as wildcard. (default: "localhost")
   --http.api value                                         API's offered over the HTTP-RPC interface (default: "eth,erigon,engine")
   --ws                                                     Enable the WS-RPC server (default: false)
   --ws.compression                                         Enable compression over WebSocket (default: false)
   --http.trace                                             Trace HTTP requests with INFO level (default: false)
   --state.cache value                                      Amount of data to store in StateCache (enabled if no --datadir set). Set 0 to disable StateCache. Defaults to 0MB (default: "0MB")
   --rpc.batch.concurrency value                            Does limit amount of goroutines to process 1 batch request. Means 1 bach request can't overload server. 1 batch still can have unlimited amount of request (default: 2)
   --rpc.streaming.disable                                  Erigon has enalbed json streaming for some heavy endpoints (like trace_*). It's treadoff: greatly reduce amount of RAM (in some cases from 30GB to 30mb), but it produce invalid json format if error happened in the middle of streaming (because json is not streaming-friendly format) (default: false)
   --db.read.concurrency value                              Does limit amount of parallel db reads. Default: equal to GOMAXPROCS (or number of CPU) (default: 80)
   --rpc.accessList value                                   Specify granular (method-by-method) API allowlist
   --trace.compat                                           Bug for bug compatibility with OE for trace_ routines (default: false)
   --rpc.gascap value                                       Sets a cap on gas that can be used in eth_call/estimateGas (default: 50000000)
   --rpc.batch.limit value                                  Maximum number of requests in a batch (default: 100)
   --rpc.returndata.limit value                             Maximum number of bytes returned from eth_call or similar invocations (default: 100000)
   --rpc.txfeecap value                                     Sets a cap on transaction fee (in ether) that can be sent via the RPC APIs (0 = no cap) (default: 1)
   --txpool.api.addr value                                  txpool api network address, for example: 127.0.0.1:9090 (default: use value of --private.api.addr)
   --trace.maxtraces value                                  Sets a limit on traces that can be returned in trace_filter (default: 200)
   --http.timeouts.read value                               Maximum duration for reading the entire request, including the body. (default: 30s)
   --http.timeouts.write value                              Maximum duration before timing out writes of the response. It is reset whenever a new request's header is read. (default: 30m0s)
   --http.timeouts.idle value                               Maximum amount of time to wait for the next request when keep-alives are enabled. If http.timeouts.idle is zero, the value of http.timeouts.read is used. (default: 2m0s)
   --authrpc.timeouts.read value                            Maximum duration for reading the entire request, including the body. (default: 30s)
   --authrpc.timeouts.write value                           Maximum duration before timing out writes of the response. It is reset whenever a new request's header is read. (default: 30m0s)
   --authrpc.timeouts.idle value                            Maximum amount of time to wait for the next request when keep-alives are enabled. If authrpc.timeouts.idle is zero, the value of authrpc.timeouts.read is used. (default: 2m0s)
   --rpc.evmtimeout value                                   Maximum amount of time to wait for the answer from EVM call. (default: 5m0s)
   --snap.keepblocks                                        Keep ancient blocks in db (useful for debug) (default: false)
   --snap.stop                                              Workaround to stop producing new snapshots, if you meet some snapshots-related critical bug. It will stop move historical data from DB to new immutable snapshots. DB will grow and may slightly slow-down - and removing this flag in future will not fix this effect (db size will not greatly reduce). (default: false)
   --db.pagesize value                                      DB is splitted to 'pages' of fixed size. Can't change DB creation. Must be power of 2 and '256b <= pagesize <= 64kb'. Default: equal to OperationSystem's pageSize. Bigger pageSize causing: 1. More writes to disk during commit 2. Smaller b-tree high 3. Less fragmentation 4. Less overhead on 'free-pages list' maintainance (a bit faster Put/Commit) 5. If expecting DB-size > 8Tb then set pageSize >= 8Kb (default: "8KB")
   --db.size.limit value                                    runtime limit of chandata db size. you can change value of this flag at any time (default: "3TB")
   --torrent.port value                                     port to listen and serve BitTorrent protocol (default: 42069)
   --torrent.maxpeers value                                 unused parameter (reserved for future use) (default: 100)
   --torrent.conns.perfile value                            connections per file (default: 10)
   --torrent.download.slots value                           amount of files to download in parallel. If network has enough seeders 1-3 slot enough, if network has lack of seeders increase to 5-7 (too big value will slow down everything). (default: 3)
   --torrent.staticpeers value                              Comma separated enode URLs to connect to
   --torrent.upload.rate value                              bytes per second, example: 32mb (default: "4mb")
   --torrent.download.rate value                            bytes per second, example: 32mb (default: "16mb")
   --torrent.verbosity value                                0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail (must set --verbosity to equal or higher level and has defeault: 3) (default: 2)
   --port value                                             Network listening port (default: 30303)
   --p2p.protocol value [ --p2p.protocol value ]            Version of eth p2p protocol (default: 68, 67)
   --p2p.allowed-ports value [ --p2p.allowed-ports value ]  Allowed ports to pick for different eth p2p protocol versions as follows <porta>,<portb>,..,<porti> (default: 30303, 30304, 30305, 30306, 30307)
   --nat value                                              NAT port mapping mechanism (any|none|upnp|pmp|stun|extip:<IP>)
                                                                 "" or "none"         default - do not nat
                                                                 "extip:77.12.33.4"   will assume the local machine is reachable on the given IP
                                                                 "any"                uses the first auto-detected mechanism
                                                                 "upnp"               uses the Universal Plug and Play protocol
                                                                 "pmp"                uses NAT-PMP with an auto-detected gateway address
                                                                 "pmp:192.168.0.1"    uses NAT-PMP with the given gateway address
                                                                 "stun"               uses STUN to detect an external IP using a default server
                                                                 "stun:<server>"      uses STUN to detect an external IP using the given server (host:port)
   --nodiscover                                             Disables the peer discovery mechanism (manual peer addition) (default: false)
   --v5disc                                                 Enables the experimental RLPx V5 (Topic Discovery) mechanism (default: false)
   --netrestrict value                                      Restricts network communication to the given IP networks (CIDR masks)
   --nodekey value                                          P2P node key file
   --nodekeyhex value                                       P2P node key as hex (for testing)
   --discovery.dns value                                    Sets DNS discovery entry points (use "" to disable DNS)
   --bootnodes value                                        Comma separated enode URLs for P2P discovery bootstrap
   --staticpeers value                                      Comma separated enode URLs to connect to
   --trustedpeers value                                     Comma separated enode URLs which are always allowed to connect, even above the peer limit
   --maxpeers value                                         Maximum number of network peers (network disabled if set to 0) (default: 100)
   --chain value                                            Name of the testnet to join (default: "mainnet")
   --dev.period value                                       Block period to use in developer mode (0 = mine only if transaction pending) (default: 0)
   --vmdebug                                                Record information useful for VM and contract debugging (default: false)
   --networkid value                                        Explicitly set network id (integer)(For testnets: use --chain <testnet_name> instead) (default: 1)
   --fakepow                                                Disables proof-of-work verification (default: false)
   --gpo.blocks value                                       Number of recent blocks to check for gas prices (default: 20)
   --gpo.percentile value                                   Suggested gas price is the given percentile of a set of recent transaction gas prices (default: 60)
   --allow-insecure-unlock                                  Allow insecure account unlocking when account-related RPCs are exposed by http (default: false)
   --experimental.history.v3                                (also known as Erigon3) Not recommended yet: Can't change this flag after node creation. New DB and Snapshots format of history allows: parallel blocks execution, get state as of given transaction without executing whole block. (default: false)
   --identity value                                         Custom node name
   --clique.checkpoint value                                number of blocks after which to save the vote snapshot to the database (default: 10)
   --clique.snapshots value                                 number of recent vote snapshots to keep in memory (default: 1024)
   --clique.signatures value                                number of recent block signatures to keep in memory (default: 16384)
   --clique.datadir value                                   a path to clique db folder
   --mine                                                   Enable mining (default: false)
   --proposer.disable                                       Disables PoS proposer (default: false)
   --miner.notify value                                     Comma separated HTTP URL list to notify of new work packages
   --miner.gaslimit value                                   Target gas limit for mined blocks (default: 30000000)
   --miner.etherbase value                                  Public address for block mining rewards (default: "0")
   --miner.extradata value                                  Block extra data set by the miner (default = client version)
   --miner.noverify                                         Disable remote sealing verification (default: false)
   --miner.sigfile value                                    Private key to sign blocks with
   --sentry.api.addr value                                  comma separated sentry addresses '<host>:<port>,<host>:<port>'
   --sentry.log-peer-info                                   Log detailed peer info when a peer connects or disconnects. Enable to integrate with observer. (default: false)
   --sentry.drop-useless-peers                              Drop useless peers, those returning empty body or header responses (default: false)
   --downloader.api.addr value                              downloader address '<host>:<port>'
   --downloader.disable.ipv4                                Turn off ipv4 for the downloader (default: false)
   --downloader.disable.ipv6                                Turns off ipv6 for the downlaoder (default: false)
   --no-downloader                                          to disable downloader component (default: false)
   --downloader.verify                                      verify snapshots on startup. it will not report founded problems but just re-download broken pieces (default: false)
   --healthcheck                                            Enable grpc health check (default: false)
   --bor.heimdall value                                     URL of Heimdall service (default: "http://localhost:1317")
   --bor.withoutheimdall                                    Run without Heimdall service (for testing purpose) (default: false)
   --bor.heimdallgRPC value                                 Address of Heimdall gRPC service
   --ethstats value                                         Reporting URL of a ethstats service (nodename:secret@host:port)
   --override.shanghaiTime value                            Manually specify Shanghai fork time, overriding the bundled setting (default: 0)
   --config value                                           Sets erigon flags from YAML/TOML file
   --lightclient.discovery.addr value                       Address for lightclient DISCV5 protocol (default: "127.0.0.1")
   --lightclient.discovery.port value                       Port for lightclient DISCV5 protocol (default: 4000)
   --lightclient.discovery.tcpport value                    TCP Port for lightclient DISCV5 protocol (default: 4001)
   --sentinel.addr value                                    Address for sentinel (default: "localhost")
   --sentinel.port value                                    Port for sentinel (default: 7777)
   --pprof                                                  Enable the pprof HTTP server (default: false)
   --pprof.addr value                                       pprof HTTP server listening interface (default: "127.0.0.1")
   --pprof.port value                                       pprof HTTP server listening port (default: 6060)
   --pprof.cpuprofile value                                 Write CPU profile to the given file
   --trace value                                            Write execution trace to the given file
   --metrics                                                Enable metrics collection and reporting (default: false)
   --metrics.addr value                                     Enable stand-alone metrics HTTP server listening interface (default: "127.0.0.1")
   --metrics.port value                                     Metrics HTTP server listening port (default: 6060)
   --log.json                                               Format console logs with JSON (default: false)
   --log.console.json                                       Format console logs with JSON (default: false)
   --log.dir.json                                           Format file logs with JSON (default: false)
   --verbosity value                                        Set the log level for console logs (default: "info")
   --log.console.verbosity value                            Set the log level for console logs (default: "info")
   --log.dir.path value                                     Path to store user and error logs to disk
   --log.dir.prefix value                                   The file name prefix for logs stored to disk
   --log.dir.verbosity value                                Set the log verbosity for logs stored to disk (default: "info")
   --help, -h                                               show help
   --version, -v                                            print the version
