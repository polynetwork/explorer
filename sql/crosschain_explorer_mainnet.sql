CREATE SCHEMA IF NOT EXISTS `cross_chain_explorer` DEFAULT CHARACTER SET utf8;
USE `cross_chain_explorer`;

DROP TABLE IF EXISTS `chain_info`;
CREATE TABLE `chain_info` (
 `xname` VARCHAR(32) NOT NULL COMMENT '链名称',
 `id`  INT(4) NOT NULL COMMENT '链id',
 `xtype` INT(4) NOT NULL COMMENT '链类型',
 `height` INT(12) NOT NULL COMMENT '解析的区块高度',
 `txin` 	INT(12) NOT NULL COMMENT '链的入金数量',
 `txout`	INT(12) NOT NULL COMMENT '链的出金数量',
 PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `chain_contract`;
CREATE TABLE `chain_contract` (
  `id` INT(4) NOT NULL COMMENT '链id',
  `contract` VARCHAR(128) NOT NULL COMMENT '跨链合约地址'
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `chain_token`;
CREATE TABLE `chain_token` (
  `id` INT(4) NOT NULL COMMENT '链id',
  `xtoken` VARCHAR(32) NOT NULL COMMENT '跨链通用token名称',
  `hash` VARCHAR(128) NOT NULL COMMENT 'token地址',
  `xname` VARCHAR(32) NOT NULL COMMENT 'token名称',
  `xtype` VARCHAR(32) NOT NULL COMMENT 'token类型',
  `xprecision` VARCHAR(32)  NOT NULL COMMENT 'token精度',
  `xdesc` VARCHAR(1024) COMMENT 'token描述'
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `chain_token_bind`;
CREATE TABLE `chain_token_bind` (
  `hash_src` VARCHAR(64) NOT NULL COMMENT '源资产',
  `hash_dest` VARCHAR(64) NOT NULL COMMENT '绑定的目标资产',
  PRIMARY KEY (`hash_src`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("poly",0,0,22732,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("btc",1,1,0,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("eth",2,2,10650091,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("ontology",3,3,9300490,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("neo",4,4,6023777,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("switcheo",5,5,202650,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("bsc",6,6,4123673,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("heco",7,7,1810758,0,0);

INSERT INTO `chain_contract`(`id`,`contract`) VALUES(0, "0300000000000000000000000000000000000000");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(2, "838bf9e95cb12dd76a54c9f9d2e3082eaf928270");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(3, "0900000000000000000000000000000000000000");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(4, "82a3401fb9a60db42c6fa2ea2b6d62e872d6257f");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ontology Gas", "0200000000000000000000000000000000000000", "ong", "OEP4", "1000000000","ong");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ontology", "0100000000000000000000000000000000000000", "ont", "OEP4", "1","ont");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ethereum", "0000000000000000000000000000000000000000", "ether", "ether", "1000000000000000000","ether");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "DeepBrain Chain", "b951ecbbc5fe37a9c280a76cb0ce0014827294cf", "DeepBrain Coin", "NEP5", "100000000","DBC");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "DeepBrain Chain", "64626331", "DeepBrain Coin", "Cosmos", "100000000","dbc1");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Switcheo", "ab38352559b8b203bde5fddfa0b07d8b2525e132", "Switcheo", "NEP5", "100000000","SWTH");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Switcheo", "73777468", "Switcheo", "Cosmos", "100000000","swth");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ontology Gas", "c8757865920e0467f5d23b59845aa357a24ea38c", "Ethereum Ontology Gas", "ERO20", "1000000000","ong");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ontology", "cb46c550539ac3db72dc7af7c89b11c306c727c2", "pONT", "erc20", "1000000000","ont");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Bitcoin", "8037dd7161401417d3571b92b86846d34309129a", "pWBTC", "OEP4", "100000000","btc");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ethereum", "a2f89e531e55636d4af1cd044237d2fd5a616c72", "oETH", "OEP4", "1000000000000000000","eth");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Tether", "3e931f60f2cd1387b52f1889dfcaf02a54b2c6a0", "oUSDT", "OEP4", "1000000","usdt");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "USD Coin", "061a07cd393aac289b8ecfda2c3784b637a2fb33", "pUSDC", "OEP4", "1000000","usdc");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Dai", "3f0def1945d7129c5f6625147dcbbaaee402e751", "oDai", "OEP4", "1000000000000000000","dai");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Tether", "dac17f958d2ee523a2206206994597c13d831ec7", "USDT", "erc20", "1000000","usdt");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "2260fac5e5542a773aa44fbcfedf7c193bc2c599", "Wrapped BTC", "erc20", "100000000","btc");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Dai", "6b175474e89094c44da98b954eedeac495271d0f", "DAI", "erc20", "1000000000000000000","dai");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "USD Coin", "a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "USD Coin", "erc20", "1000000","usdc");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ontology", "33ae7eae016193ba0fe238b223623bc78faac158", "ontd", "oep4", "1000000000","ont");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "eb4c2781e4eba804ce9a9803c67d0893436bb27d", "renBTC", "ERC20", "100000000","renBTC");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Bitcoin", "46c3051c553aaeb3724ea69336ec483f39fa91b1", "prenBTC", "OEP4", "100000000","renBTC");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Wing", "db0f18081b505a7de20b18ac41856bcb4ba86a1a", "pWING", "ERC20", "1000000000","Wing");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Wing", "00c59fcd27a562d6397883eab1f2fff56e58ef80", "Wing Token (WING)", "OEP4", "1000000000","Wing");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "sUSD", "57ab1ec28d129707052df4df418d58a2d46d5f51", "Synth sUSD", "ERC20", "1000000000000000000","sUSD");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "sUSD", "17a58a4a65959c2f567e5063c560f9d09fb81284", "psUSD", "OEP4", "1000000000000000000","sUSD");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Paxos Standard", "8e870d67f660d95d5be530380d0ec0bd388289e1", "Paxos Standard", "ERC20", "1000000000000000000","PAX");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Paxos Standard", "0dabee6055a1c17e3b4bcb15af1a713605b7fcfc", "pPAX", "OEP4", "1000000000000000000","PAX");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Neo", "381225768DD2bd60D70482B51109D0DEFeE92503", "Poly NEO Token", "ERC20", "1000000000","pNEO");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Neo", "9a576d927dda934b8ce69f35ec2c1025ceb10e6f", "pNEO", "OEP4", "1000000000","pNEO");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Ethereum", "17c76859c11bc14da5b3e9c88fa695513442c606", "nETH", "NEP5", "1000000000000000000","Ethereum");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Ontology", "271e1e4616158c7440ffd1d5ca51c0c12c792833", "nONT", "NEP5", "1000000000","Ontology");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Ethereum", "0df563008be710f3e0130208f8adc95ed7e5518d", "pnWETH", "NEP5", "1000000000000","WETH");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ethereum", "e179198fd42f5de1a04ffd9a36d6dc428ceb13f7", "nWETH", "ERC20", "1000000000000","WETH");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Neo", "f46719e2d16bf50cddcef9d4bbfece901f73cbb6", "nNEO", "NEP5", "100000000","nNEO");





INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Ontology", "c277117879af3197fbef92c71e95800aa3b89d9a", "pONT", "NEP5", "1000000000","ONT");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Tether", "282e3340d5a1cd6a461d5f558d91bc1dbc02a07b", "pnUSDT", "NEP5", "1000000","USDT");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Tether", "2205d2f559ef91580090011aa4e0ef68ec33da44", "nUSDT", "ERC20", "1000000","USDT");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Bitcoin", "534dcac35b0dfadc7b2d716a7a73a7067c148b37", "pnWBTC", "NEP5", "100000000","pnWBTC");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "bb44b36e588445d7da61a1e2e426664d03d40888", "nWBTC", "ERC20", "100000000","nWBTC");





INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Dai", "7b956c0c11fcffb9c9227ca1925ba4c3486b36f1", "pDAI", "OEP4", "1000000000000000000","DAI");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ethereum", "df19600d334bb13c6a9e3e9777aa8ec6ed6a4a79", "pETH", "OEP4", "1000000000000000000","pETH");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Tether", "ac654837a90eee8fccabd87a2d4fc7637484f01a", "pUSDT", "OEP4", "1000000","pUSDT");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "7757ffe3ac09bc6430f6896f720e77cf80ec1f74", "nrenBTC", "ERC20", "100000000","nrenBTC");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "2dd56dc238d1fc2f9aac3793a287f4e0af1b08b4", "nsBTC", "ERC20", "1000000000000000000","nsBTC");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "TrueUSD", "886f6F287Bb2eA7DE03830a5FD339EDc107c559f", "nTUSD", "ERC20", "1000000000000000000","nTUSD");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "USD Coin", "7f0ad0525cb8c17d3f5c06ceb0aea20fa0d2ca0a", "nUSDC", "ERC20", "1000000","nUSDC");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Dai", "7245ded8459f59b0a680640535476c11b3cd0ef6", "nDAI", "ERC20", "1000000000000000000","nDAI");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "UNI_V2_ETH_WBTC", "6c5fa7a3c2cd98a689b1305bd38b56120fe15744", "normalized Uniswap V2", "ERC20", "1000000000000000000","nUNI-V2");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "UNI_V2_ETH_WBTC", "c534d65c85c074887f58ed1f3bad7dfd739a525e", "pnUNI_V2_ETH_WBTC", "NEP5", "1000000000000000000","nUNI-V2");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "FLM", "c6061ca95ad0378bdb12381206a1d723d14b72c4", "Poly Flamingo Token", "ERC20", "100000000","FLM");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "FLM", "4d9eab13620fe3569ba3b0e56e2877739e4145e3", "Flamingo", "NEP5", "100000000","FLM");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Uniswap", "1f9840a85d5af5bf1d1762f925bdaddc4201f984", "Uniswap", "ERC20", "1000000000000000000","Uniswap");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Uniswap", "89029ef258b82c5c3741fe25db91375e9301dc71", "pUNI", "OEP4", "1000000000000000000","Uniswap");


INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("0200000000000000000000000000000000000000", "0200000000000000000000000000000000000000");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("0100000000000000000000000000000000000000", "0100000000000000000000000000000000000000");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("0000000000000000000000000000000000000000", "0000000000000000000000000000000000000000");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("b951ecbbc5fe37a9c280a76cb0ce0014827294cf", "b951ecbbc5fe37a9c280a76cb0ce0014827294cf");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("64626331", "b951ecbbc5fe37a9c280a76cb0ce0014827294cf");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("ab38352559b8b203bde5fddfa0b07d8b2525e132", "ab38352559b8b203bde5fddfa0b07d8b2525e132");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("73777468", "ab38352559b8b203bde5fddfa0b07d8b2525e132");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("c8757865920e0467f5d23b59845aa357a24ea38c", "0200000000000000000000000000000000000000");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("cb46c550539ac3db72dc7af7c89b11c306c727c2", "33ae7eae016193ba0fe238b223623bc78faac158");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("cb46c550539ac3db72dc7af7c89b11c306c727c2", "cb46c550539ac3db72dc7af7c89b11c306c727c2");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("8037dd7161401417d3571b92b86846d34309129a", "2260fac5e5542a773aa44fbcfedf7c193bc2c599");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("a2f89e531e55636d4af1cd044237d2fd5a616c72", "0000000000000000000000000000000000000000");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("3e931f60f2cd1387b52f1889dfcaf02a54b2c6a0", "dac17f958d2ee523a2206206994597c13d831ec7");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("061a07cd393aac289b8ecfda2c3784b637a2fb33", "a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("3f0def1945d7129c5f6625147dcbbaaee402e751", "6b175474e89094c44da98b954eedeac495271d0f");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("dac17f958d2ee523a2206206994597c13d831ec7", "dac17f958d2ee523a2206206994597c13d831ec7");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("2260fac5e5542a773aa44fbcfedf7c193bc2c599", "2260fac5e5542a773aa44fbcfedf7c193bc2c599");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("6b175474e89094c44da98b954eedeac495271d0f", "6b175474e89094c44da98b954eedeac495271d0f");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("33ae7eae016193ba0fe238b223623bc78faac158", "33ae7eae016193ba0fe238b223623bc78faac158");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("eb4c2781e4eba804ce9a9803c67d0893436bb27d", "eb4c2781e4eba804ce9a9803c67d0893436bb27d");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("46c3051c553aaeb3724ea69336ec483f39fa91b1", "eb4c2781e4eba804ce9a9803c67d0893436bb27d");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("6a4c89eb9a26a2da34f13f8976daa9fd7526f35c", "0200000000000000000000000000000000000000");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("db0f18081b505a7de20b18ac41856bcb4ba86a1a", "00c59fcd27a562d6397883eab1f2fff56e58ef80");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("00c59fcd27a562d6397883eab1f2fff56e58ef80", "00c59fcd27a562d6397883eab1f2fff56e58ef80");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("57ab1ec28d129707052df4df418d58a2d46d5f51", "57ab1ec28d129707052df4df418d58a2d46d5f51");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("17a58a4a65959c2f567e5063c560f9d09fb81284", "57ab1ec28d129707052df4df418d58a2d46d5f51");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("8e870d67f660d95d5be530380d0ec0bd388289e1", "8e870d67f660d95d5be530380d0ec0bd388289e1");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("0dabee6055a1c17e3b4bcb15af1a713605b7fcfc", "8e870d67f660d95d5be530380d0ec0bd388289e1");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("381225768dd2bd60d70482b51109d0defee92503", "f46719e2d16bf50cddcef9d4bbfece901f73cbb6");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("9a576d927dda934b8ce69f35ec2c1025ceb10e6f", "f46719e2d16bf50cddcef9d4bbfece901f73cbb6");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("17c76859c11bc14da5b3e9c88fa695513442c606", "0000000000000000000000000000000000000000");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("271e1e4616158c7440ffd1d5ca51c0c12c792833", "33ae7eae016193ba0fe238b223623bc78faac158");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("0df563008be710f3e0130208f8adc95ed7e5518d", "e179198fd42f5de1a04ffd9a36d6dc428ceb13f7");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("e179198fd42f5de1a04ffd9a36d6dc428ceb13f7", "e179198fd42f5de1a04ffd9a36d6dc428ceb13f7");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("f46719e2d16bf50cddcef9d4bbfece901f73cbb6", "f46719e2d16bf50cddcef9d4bbfece901f73cbb6");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("c277117879af3197fbef92c71e95800aa3b89d9a", "33ae7eae016193ba0fe238b223623bc78faac158");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("282e3340d5a1cd6a461d5f558d91bc1dbc02a07b", "2205d2f559ef91580090011aa4e0ef68ec33da44");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("2205d2f559ef91580090011aa4e0ef68ec33da44", "2205d2f559ef91580090011aa4e0ef68ec33da44");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("534dcac35b0dfadc7b2d716a7a73a7067c148b37", "bb44b36e588445d7da61a1e2e426664d03d40888");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("bb44b36e588445d7da61a1e2e426664d03d40888", "bb44b36e588445d7da61a1e2e426664d03d40888");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("7757ffe3ac09bc6430f6896f720e77cf80ec1f74", "7757ffe3ac09bc6430f6896f720e77cf80ec1f74");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("2dd56dc238d1fc2f9aac3793a287f4e0af1b08b4", "2dd56dc238d1fc2f9aac3793a287f4e0af1b08b4");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("886f6f287bb2ea7de03830a5fd339edc107c559f", "886f6f287bb2ea7de03830a5fd339edc107c559f");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("7f0ad0525cb8c17d3f5c06ceb0aea20fa0d2ca0a", "7f0ad0525cb8c17d3f5c06ceb0aea20fa0d2ca0a");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("7245ded8459f59b0a680640535476c11b3cd0ef6", "7245ded8459f59b0a680640535476c11b3cd0ef6");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("7b956c0c11fcffb9c9227ca1925ba4c3486b36f1", "6b175474e89094c44da98b954eedeac495271d0f");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("df19600d334bb13c6a9e3e9777aa8ec6ed6a4a79", "0000000000000000000000000000000000000000");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("ac654837a90eee8fccabd87a2d4fc7637484f01a", "dac17f958d2ee523a2206206994597c13d831ec7");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("6c5fa7a3c2cd98a689b1305bd38b56120fe15744", "6c5fa7a3c2cd98a689b1305bd38b56120fe15744");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("c534d65c85c074887f58ed1f3bad7dfd739a525e", "6c5fa7a3c2cd98a689b1305bd38b56120fe15744");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("c6061ca95ad0378bdb12381206a1d723d14b72c4", "4d9eab13620fe3569ba3b0e56e2877739e4145e3");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("4d9eab13620fe3569ba3b0e56e2877739e4145e3", "4d9eab13620fe3569ba3b0e56e2877739e4145e3");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("1f9840a85d5af5bf1d1762f925bdaddc4201f984", "1f9840a85d5af5bf1d1762f925bdaddc4201f984");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("89029ef258b82c5c3741fe25db91375e9301dc71", "1f9840a85d5af5bf1d1762f925bdaddc4201f984");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "yearn.finance", "0bc529c00c6401aef6d220be8c6ea1667f6ad93e", "YFI", "ERC20", "1000000000000000000","yearn.finance");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "yearn.finance", "8c8dc28d10c22a66357240b920b569e32447d7af", "pYFI", "OEP4", "1000000000000000000","pYFI");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Unifi Protocol DAO", "441761326490cacf7af299725b6292597ee822c2", "UNFI", "ERC20", "1000000000000000000","UNFI");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Unifi Protocol DAO", "6f560d392a8701d0931a7d61f8ac4bdcc050e9ab", "pUNFI", "OEP4", "1000000000000000000","pUNFI");


INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("0bc529c00c6401aef6d220be8c6ea1667f6ad93e", "0bc529c00c6401aef6d220be8c6ea1667f6ad93e");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("8c8dc28d10c22a66357240b920b569e32447d7af", "0bc529c00c6401aef6d220be8c6ea1667f6ad93e");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("441761326490cacf7af299725b6292597ee822c2", "441761326490cacf7af299725b6292597ee822c2");
INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("6f560d392a8701d0931a7d61f8ac4bdcc050e9ab", "441761326490cacf7af299725b6292597ee822c2");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(6, "Neo", "b119b3b8e5e6eeffbe754b20ee5b8a42809931fb", "pNEO", "ERC20", "1000000000","pNEO");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("b119b3b8e5e6eeffbe754b20ee5b8a42809931fb", "f46719e2d16bf50cddcef9d4bbfece901f73cbb6");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(6, "Ethereum", "b9478391eec218defa96f7b9a7938cf44e7a2fd5", "pETH", "ERC20", "1000000000000000000","pETH");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("b9478391eec218defa96f7b9a7938cf44e7a2fd5", "0000000000000000000000000000000000000000");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(6, "Tether", "48389753b64c9e581975457332e60dc49325a653", "pUSDT", "ERC20", "1000000","pUSDT");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("48389753b64c9e581975457332e60dc49325a653", "dac17f958d2ee523a2206206994597c13d831ec7");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(6, "Dai", "8f339abc2a2a8a4d0364c7e35f892c40fbfb4bc0", "pDAI", "ERC20", "1000000000000000000","pDAI");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("8f339abc2a2a8a4d0364c7e35f892c40fbfb4bc0", "6b175474e89094c44da98b954eedeac495271d0f");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(6, "USD Coin", "0dbbf67fb78651d3f6407a421040f1503b486693", "pUSDC", "ERC20", "1000000","pUSDC");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("0dbbf67fb78651d3f6407a421040f1503b486693", "a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(6, "sUSD", "89bcd91f7922126c568436841b16d036528e9714", "psUSD", "ERC20", "1000000000000000000","psUSD");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("89bcd91f7922126c568436841b16d036528e9714", "57ab1ec28d129707052df4df418d58a2d46d5f51");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "Neo", "6514a5ebff7944099591ae3e8a5c0979c83b2571", "pNEO", "ERC20", "1000000000","pNEO");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("6514a5ebff7944099591ae3e8a5c0979c83b2571", "f46719e2d16bf50cddcef9d4bbfece901f73cbb6");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "Ethereum", "8c0859c191d8f100e4a3c0d8c0066c36a0c1f894", "pETH", "ERC20", "1000000000000000000","pETH");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("8c0859c191d8f100e4a3c0d8c0066c36a0c1f894", "0000000000000000000000000000000000000000");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "Tether", "a7d1aac3c9bf61559c25f94132a9f801e8b5f97e", "pUSDT", "ERC20", "1000000","pUSDT");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("a7d1aac3c9bf61559c25f94132a9f801e8b5f97e", "dac17f958d2ee523a2206206994597c13d831ec7");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "Dai", "643f3914fb8ede03d932c79732746a8c11ae470a", "pDAI", "ERC20", "1000000000000000000","pDAI");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("643f3914fb8ede03d932c79732746a8c11ae470a", "6b175474e89094c44da98b954eedeac495271d0f");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "USD Coin", "e85631b817923487ba40263144eeff532cff10a2", "pUSDC", "ERC20", "1000000","pUSDC");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("e85631b817923487ba40263144eeff532cff10a2", "a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "sUSD", "002e47d940dfd177dc0fe78321e17ef84676985d", "psUSD", "ERC20", "1000000000000000000","psUSD");

INSERT INTO `chain_token_bind`(`hash_src`,`hash_dest`) VALUES("002e47d940dfd177dc0fe78321e17ef84676985d", "57ab1ec28d129707052df4df418d58a2d46d5f51");




DROP TABLE IF EXISTS `poly_validators`;
CREATE TABLE `poly_validators` (
  `height` INT(12) NOT NULL COMMENT '交易的高度',
  `validators`  VARCHAR(8192) COMMENT '验证节点',
  PRIMARY KEY (`height`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `mchain_tx`;
CREATE TABLE `mchain_tx` (
 `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
 `chain_id` INT(4) NOT NULL COMMENT '链ID',
 `state` INT(1) NOT NULL COMMENT '交易状态',
 `tt` INT(11) NOT NULL COMMENT '交易时间',
 `fee` BIGINT(8) NOT NULL COMMENT '交易手续费',
 `height` INT(12) NOT NULL COMMENT '交易的高度',
 `fchain` INT(4) NOT NULL COMMENT '源链的id',
 `ftxhash` VARCHAR(128) NOT NULL COMMENT '源链的交易hash',
 `tchain` INT(4) NOT NULL COMMENT '目标链的id',
 `xkey` VARCHAR(8192) COMMENT '比特币交易',
 PRIMARY KEY (`txhash`),
 UNIQUE (`ftxhash`),
 INDEX (`tt`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `fchain_tx`;
CREATE TABLE `fchain_tx` (
 `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
 `chain_id` INT(4) NOT NULL COMMENT '链ID',
 `state` INT(1) NOT NULL COMMENT '交易状态',
 `tt` INT(4) NOT NULL COMMENT '交易时间',
 `fee` BIGINT(8) NOT NULL COMMENT '交易手续费',
 `height` INT(12) NOT NULL COMMENT '交易的高度',
 `xuser` VARCHAR(128) NOT NULL COMMENT '用户',
 `tchain` INT(4) NOT NULL COMMENT '目标链的id',
 `contract` VARCHAR(128) NOT NULL COMMENT '执行的合约',
 `xkey` VARCHAR(8192) NOT NULL COMMENT '目标链的参数',
 `xparam` VARCHAR(8192) NOT NULL COMMENT '合约参数',
 PRIMARY KEY (`txhash`),
 INDEX (`tt`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `fchain_transfer`;
CREATE TABLE `fchain_transfer` (
  `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
  `chain_id` INT(4) NOT NULL COMMENT '链ID',
  `tt` INT(4) NOT NULL COMMENT '交易时间',
  `asset` VARCHAR(128) NOT NULL COMMENT '资产hash',
  `xfrom` VARCHAR(128) NOT NULL COMMENT '发送用户',
  `xto` VARCHAR(128) NOT NULL COMMENT '接受用户',
  `amount` VARCHAR(64) NOT NULL COMMENT '收到的金额',
  `tochainid` INT(4) NOT NULL COMMENT '目标链的id',
  `toasset` VARCHAR(1024) NOT NULL COMMENT '目标链的资产hash',
  `touser` VARCHAR(128) NOT NULL COMMENT '目标链的接受用户',
  PRIMARY KEY (`txhash`),
  INDEX (`asset`, `tt`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tchain_tx`;
CREATE TABLE `tchain_tx` (
 `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
 `chain_id` INT(4) NOT NULL COMMENT '链ID',
 `state` INT(1) NOT NULL COMMENT '交易状态',
 `tt` INT(4) NOT NULL COMMENT '交易时间',
 `fee` BIGINT(8) NOT NULL COMMENT '交易手续费',
 `height` INT(12) NOT NULL COMMENT '交易的高度',
 `fchain` INT(4) NOT NULL COMMENT '源链的id',
 `contract` VARCHAR(128) NOT NULL COMMENT '执行的合约',
 `rtxhash` VARCHAR(128) NOT NULL COMMENT '中继链的交易hash',
 PRIMARY KEY (`txhash`),
 UNIQUE (`rtxhash`),
 INDEX (`tt`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tchain_transfer`;
CREATE TABLE `tchain_transfer` (
  `txhash`  VARCHAR(128) NOT NULL COMMENT '交易hash',
  `chain_id` INT(4) NOT NULL COMMENT '链ID',
  `tt` INT(4) NOT NULL COMMENT '交易时间',
  `asset` VARCHAR(128) NOT NULL COMMENT '资产hash',
  `xfrom` VARCHAR(128) NOT NULL COMMENT '发送用户',
  `xto` VARCHAR(128) NOT NULL COMMENT '接受用户',
  `amount` VARCHAR(64) NOT NULL COMMENT '收到的金额',
  PRIMARY KEY (`txhash`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `asset_statistic`;
CREATE TABLE `asset_statistic` (
  `xname` VARCHAR(16)  COMMENT '资产名称',
  `addressnum`   INT(4) NOT NULL COMMENT '资产的总地址数',
  `amount`       BIGINT(8)  NOT NULL COMMENT '资产的总价值',
  `amount_btc`  BIGINT(8)  NOT NULL COMMENT '资产的总价值',
  `amount_usd`  BIGINT(8)  NOT NULL COMMENT '资产的总价值',
  `txnum`       INT(4) NOT NULL COMMENT '总的交易个数',
  `latestupdate` INT(4)  NOT NULL COMMENT '统计数据的时间点',
  PRIMARY KEY (`xname`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `transfer_statistic`;
CREATE TABLE `transfer_statistic` (
  `asset`        VARCHAR(64) COMMENT '资产hash',
  `amount`       BIGINT(8)  NOT NULL COMMENT '资产总额',
  `latestin`     INT(4)  NOT NULL COMMENT '统计数据的时间点',
  `latestout`    INT(4)  NOT NULL COMMENT '统计数据的时间点',
  PRIMARY KEY (`asset`)
)ENGINE=INNODB DEFAULT CHARSET=utf8;

SET sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));
