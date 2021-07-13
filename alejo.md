## latest Janus fix

Replaces `gettransaction` with `getrawtransaction` in file eth_getTransactionReceipt.go

## theGraph logs

```bash
2021-07-06T20:42:03.652390581Z Jul 06 20:42:03.652 INFO Graph Node version: 0.22.0 (2021-02-24)
2021-07-06T20:42:03.652415838Z Jul 06 20:42:03.652 INFO Generating configuration from command line arguments
2021-07-06T20:42:03.668006325Z Jul 06 20:42:03.667 INFO Starting up
2021-07-06T20:42:03.668016490Z Jul 06 20:42:03.667 INFO Trying IPFS node at: http://ipfs:5001/
2021-07-06T20:42:03.675032400Z Jul 06 20:42:03.674 INFO Creating transport, capabilities: archive, trace, url: http://172.25.0.1:23889, network: mainnet
2021-07-06T20:42:03.676722184Z Jul 06 20:42:03.676 INFO Successfully connected to IPFS node at: http://ipfs:5001/
2021-07-06T20:42:03.693811901Z Jul 06 20:42:03.693 INFO Connecting to Postgres, weight: 1, conn_pool_size: 10, url: postgresql://graph-node:HIDDEN_PASSWORD@postgres:5432/graph-node, pool: main, shard: primary
2021-07-06T20:42:03.706556502Z Jul 06 20:42:03.706 INFO Pool successfully connected to Postgres, pool: main, shard: primary, component: Store
2021-07-06T20:42:03.710187934Z Jul 06 20:42:03.710 INFO Waiting for other graph-node instances to finish migrating, shard: primary, component: Store
2021-07-06T20:42:03.710610698Z Jul 06 20:42:03.710 INFO Running migrations, shard: primary, component: Store
2021-07-06T20:42:04.426583297Z Jul 06 20:42:04.426 INFO Migrations finished, shard: primary, component: Store
2021-07-06T20:42:04.426979279Z Jul 06 20:42:04.426 INFO Connecting to Ethereum..., capabilities: archive, trace, network: mainnet
2021-07-06T20:42:04.444631892Z Jul 06 20:42:04.444 INFO Connected to Ethereum, capabilities: archive, trace, network_version: 81, network: mainnet
2021-07-06T20:42:04.471130417Z Jul 06 20:42:04.471 INFO Creating LoadManager in disabled mode, component: LoadManager
2021-07-06T20:42:04.471151957Z Jul 06 20:42:04.471 INFO Starting block ingestors
2021-07-06T20:42:04.471156342Z Jul 06 20:42:04.471 INFO Starting block ingestor for network, network_name: mainnet
2021-07-06T20:42:04.472181666Z Jul 06 20:42:04.472 INFO Starting JSON-RPC admin server at: http://localhost:8020, component: JsonRpcServer
2021-07-06T20:42:04.472249962Z Jul 06 20:42:04.472 INFO Started all subgraphs, component: SubgraphRegistrar
2021-07-06T20:42:04.472437811Z Jul 06 20:42:04.472 INFO Starting GraphQL HTTP server at: http://localhost:8000, component: GraphQLServer
2021-07-06T20:42:04.472444786Z Jul 06 20:42:04.472 INFO Starting index node server at: http://localhost:8030, component: IndexNodeServer
2021-07-06T20:42:04.472448377Z Jul 06 20:42:04.472 INFO Starting metrics server at: http://localhost:8040, component: MetricsServer
2021-07-06T20:42:04.472451936Z Jul 06 20:42:04.472 INFO Starting GraphQL WebSocket server at: ws://localhost:8001, component: SubscriptionServer
2021-07-06T20:42:04.851234328Z Jul 06 20:42:04.544 INFO Downloading latest blocks from Ethereum. This may take a few minutes..., network_name: mainnet, component: BlockIngestor
2021-07-06T20:42:12.245015935Z Jul 06 20:42:12.244 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3175, latest_block_head: 545105, current_block_head: 541930, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:42:19.004048799Z Jul 06 20:42:19.003 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2754, latest_block_head: 547859, current_block_head: 545105, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:42:25.914727332Z Jul 06 20:42:25.914 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2846, latest_block_head: 550705, current_block_head: 547859, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:42:33.396605250Z Jul 06 20:42:33.396 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2933, latest_block_head: 553638, current_block_head: 550705, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:42:41.148440001Z Jul 06 20:42:41.041 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2960, latest_block_head: 556598, current_block_head: 553638, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:42:48.082815791Z Jul 06 20:42:48.082 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2904, latest_block_head: 559502, current_block_head: 556598, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:42:55.459169377Z Jul 06 20:42:55.458 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3119, latest_block_head: 562621, current_block_head: 559502, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:43:03.576187956Z Jul 06 20:43:03.574 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3171, latest_block_head: 565792, current_block_head: 562621, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:43:10.069670815Z Jul 06 20:43:10.069 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2829, latest_block_head: 568621, current_block_head: 565792, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:43:15.580756401Z Jul 06 20:43:15.580 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2172, latest_block_head: 570793, current_block_head: 568621, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:43:22.088444711Z Jul 06 20:43:22.088 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2786, latest_block_head: 573579, current_block_head: 570793, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:43:28.606368506Z Jul 06 20:43:28.606 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2715, latest_block_head: 576294, current_block_head: 573579, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:43:34.850466749Z Jul 06 20:43:34.850 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2391, latest_block_head: 578685, current_block_head: 576294, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:43:41.751683736Z Jul 06 20:43:41.751 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2865, latest_block_head: 581550, current_block_head: 578685, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:43:48.272208954Z Jul 06 20:43:48.272 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2745, latest_block_head: 584295, current_block_head: 581550, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:43:54.941983267Z Jul 06 20:43:54.941 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2672, latest_block_head: 586967, current_block_head: 584295, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:44:02.905250926Z Jul 06 20:44:02.905 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2847, latest_block_head: 589814, current_block_head: 586967, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:44:11.913527346Z Jul 06 20:44:11.913 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3459, latest_block_head: 593273, current_block_head: 589814, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:44:20.918856280Z Jul 06 20:44:20.918 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3271, latest_block_head: 596544, current_block_head: 593273, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:44:27.787188962Z Jul 06 20:44:27.787 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2827, latest_block_head: 599371, current_block_head: 596544, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:44:34.682759307Z Jul 06 20:44:34.679 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2663, latest_block_head: 602034, current_block_head: 599371, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:44:41.832176399Z Jul 06 20:44:41.832 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2858, latest_block_head: 604892, current_block_head: 602034, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:44:50.871681604Z Jul 06 20:44:50.871 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3320, latest_block_head: 608212, current_block_head: 604892, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:44:58.467056020Z Jul 06 20:44:58.466 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2800, latest_block_head: 611012, current_block_head: 608212, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:45:06.654594575Z Jul 06 20:45:06.654 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3309, latest_block_head: 614321, current_block_head: 611012, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:45:20.631925091Z Jul 06 20:45:20.631 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3989, latest_block_head: 618310, current_block_head: 614321, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:45:29.349779188Z Jul 06 20:45:29.349 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3382, latest_block_head: 621692, current_block_head: 618310, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:45:39.074278260Z Jul 06 20:45:39.074 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3773, latest_block_head: 625465, current_block_head: 621692, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:45:46.192174523Z Jul 06 20:45:46.192 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2503, latest_block_head: 627968, current_block_head: 625465, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:45:53.993828272Z Jul 06 20:45:53.993 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2660, latest_block_head: 630628, current_block_head: 627968, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:46:00.443208853Z Jul 06 20:46:00.443 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2254, latest_block_head: 632882, current_block_head: 630628, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:46:07.613763157Z Jul 06 20:46:07.613 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2489, latest_block_head: 635371, current_block_head: 632882, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:46:13.745277997Z Jul 06 20:46:13.745 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2163, latest_block_head: 637534, current_block_head: 635371, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:46:20.190721647Z Jul 06 20:46:20.190 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2250, latest_block_head: 639784, current_block_head: 637534, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:46:26.764886262Z Jul 06 20:46:26.764 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2098, latest_block_head: 641882, current_block_head: 639784, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:46:36.989810919Z Jul 06 20:46:36.989 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3496, latest_block_head: 645378, current_block_head: 641882, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:46:44.195152047Z Jul 06 20:46:44.195 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2355, latest_block_head: 647733, current_block_head: 645378, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:46:52.003315813Z Jul 06 20:46:52.003 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2482, latest_block_head: 650215, current_block_head: 647733, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:47:04.642419837Z Jul 06 20:47:04.642 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 3810, latest_block_head: 654025, current_block_head: 650215, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:47:12.800167437Z Jul 06 20:47:12.800 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2438, latest_block_head: 656463, current_block_head: 654025, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:47:22.028071586Z Jul 06 20:47:22.027 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2699, latest_block_head: 659162, current_block_head: 656463, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:47:31.074595685Z Jul 06 20:47:31.074 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2706, latest_block_head: 661868, current_block_head: 659162, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:47:39.440665928Z Jul 06 20:47:39.440 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2507, latest_block_head: 664375, current_block_head: 661868, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:47:47.705807960Z Jul 06 20:47:47.705 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2520, latest_block_head: 666895, current_block_head: 664375, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:47:55.199826992Z Jul 06 20:47:55.199 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2154, latest_block_head: 669049, current_block_head: 666895, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:48:02.011964270Z Jul 06 20:48:02.011 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 1891, latest_block_head: 670940, current_block_head: 669049, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:48:09.954375299Z Jul 06 20:48:09.954 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2376, latest_block_head: 673316, current_block_head: 670940, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:48:18.259157751Z Jul 06 20:48:18.259 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2501, latest_block_head: 675817, current_block_head: 673316, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:48:27.076653057Z Jul 06 20:48:27.076 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2716, latest_block_head: 678533, current_block_head: 675817, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:48:36.166405277Z Jul 06 20:48:36.166 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2756, latest_block_head: 681289, current_block_head: 678533, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:48:45.419714917Z Jul 06 20:48:45.419 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2688, latest_block_head: 683977, current_block_head: 681289, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:48:54.931096783Z Jul 06 20:48:54.930 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2791, latest_block_head: 686768, current_block_head: 683977, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:49:03.280615029Z Jul 06 20:49:03.280 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2468, latest_block_head: 689236, current_block_head: 686768, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:49:12.526391643Z Jul 06 20:49:12.526 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2613, latest_block_head: 691849, current_block_head: 689236, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:49:19.970475216Z Jul 06 20:49:19.970 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 1983, latest_block_head: 693832, current_block_head: 691849, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:49:28.426519647Z Jul 06 20:49:28.426 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2501, latest_block_head: 696333, current_block_head: 693832, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:49:38.794603860Z Jul 06 20:49:38.794 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2902, latest_block_head: 699235, current_block_head: 696333, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:49:46.866403990Z Jul 06 20:49:46.866 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2253, latest_block_head: 701488, current_block_head: 699235, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:49:57.380416716Z Jul 06 20:49:57.380 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2717, latest_block_head: 704205, current_block_head: 701488, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:50:04.302275278Z Jul 06 20:50:04.302 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 1424, latest_block_head: 705629, current_block_head: 704205, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:52:38.111417083Z Jul 06 20:52:38.111 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 44301, latest_block_head: 749930, current_block_head: 705629, network_name: mainnet, component: Blo
ckIngestor
2021-07-06T20:52:45.558023332Z Jul 06 20:52:45.557 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2347, latest_block_head: 752277, current_block_head: 749930, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:52:51.934394214Z Jul 06 20:52:51.934 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 1965, latest_block_head: 754242, current_block_head: 752277, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:52:59.244986941Z Jul 06 20:52:59.244 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2344, latest_block_head: 756586, current_block_head: 754242, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:53:06.816290889Z Jul 06 20:53:06.816 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2353, latest_block_head: 758939, current_block_head: 756586, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:53:15.277407901Z Jul 06 20:53:15.277 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2614, latest_block_head: 761553, current_block_head: 758939, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:53:22.596649306Z Jul 06 20:53:22.596 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2341, latest_block_head: 763894, current_block_head: 761553, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:53:29.891015019Z Jul 06 20:53:29.890 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2299, latest_block_head: 766193, current_block_head: 763894, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:53:38.861412830Z Jul 06 20:53:38.861 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2780, latest_block_head: 768973, current_block_head: 766193, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:53:46.815413359Z Jul 06 20:53:46.815 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2489, latest_block_head: 771462, current_block_head: 768973, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:53:53.487725256Z Jul 06 20:53:53.487 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 1800, latest_block_head: 773262, current_block_head: 771462, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:54:00.572728822Z Jul 06 20:54:00.572 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 1781, latest_block_head: 775043, current_block_head: 773262, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:54:03.506609350Z Jul 06 20:54:03.506 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 254, latest_block_head: 775297, current_block_head: 775043, network_name: mainnet, component: Block
Ingestor
2021-07-06T20:54:06.990799288Z Jul 06 20:54:06.990 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 527, latest_block_head: 775824, current_block_head: 775297, network_name: mainnet, component: Block
Ingestor
2021-07-06T20:54:28.182440160Z Jul 06 20:54:28.182 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2723, latest_block_head: 778547, current_block_head: 775824, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:54:35.365910277Z Jul 06 20:54:35.365 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 1915, latest_block_head: 780462, current_block_head: 778547, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:54:43.673146715Z Jul 06 20:54:43.672 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2531, latest_block_head: 782993, current_block_head: 780462, network_name: mainnet, component: Bloc
kIngestor
2021-07-06T20:54:50.602537590Z Jul 06 20:54:50.602 INFO Syncing 50 blocks from Ethereum., code: BlockIngestionLagging, blocks_needed: 50, blocks_behind: 2215, latest_block_head: 785208, current_block_head: 782993, network_name: mainnet, component: Bloc
kIngestor
```

## updated Janus and theGraph logs

https://docs.google.com/spreadsheets/d/1gi9B8VnIFgbhwmZbc4BH5PX4miyS_LZphyyR5nwEmg4/edit?usp=sharing

