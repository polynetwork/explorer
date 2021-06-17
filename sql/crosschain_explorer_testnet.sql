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

INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("poly",0,0,276644,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("btc",1,1,0,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("eth",2,2,8631110,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("ontology",3,3,13653290,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("neo",5,5,5487332,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("cosmos",5,5,0,0,0);

INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("heco",7,7,1378216,0,0);

INSERT INTO `chain_contract`(`id`,`contract`) VALUES(0, "0300000000000000000000000000000000000000");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(2, "726532586c50ec9f4080b71f906a3d9779bbd64f");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(3, "0900000000000000000000000000000000000000");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(5, "e1695b1314a1331e3935481620417ed835669407");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(1, "Bitcoin", "0000000000000000000000000000000000000011", "btc", "BTC", "100000000","btc");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "92705a16815a3d1aec3ce9cc273c5aa302961fcc", "BTC Token (BTCX)", "ERC20", "100000000","btc");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Bitcoin", "3ee29d5cc82771e91383f9ba09c6f5c5878f3f24", "btcx", "NEP5", "100000000","btc");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Bitcoin", "62746378", "btc", "Cosmos", "100000000","btc");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "oep4", "3105a14f7956d33a51f12ef3ae50a3f1ef161dff", "OEP4 Template (OEP4T)", "ERC20", "1000000000","oep4");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "oep4", "969850e009b5e2a061694f3479ec8e44bc68bcd3", "OEP4Template", "OEP4", "1000000000","oep4");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "oep4", "6f65703478", "oep4", "Cosmos", "1000000000","oep4");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ontology Gas", "42d9fef0cbd9c3000cece9764d99a4a6fe9e1b34", "ONG Token (ONGX)", "ERC20", "1000000000","ong");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ontology Gas", "0200000000000000000000000000000000000000", "ong", "OEP4", "1000000000","ong");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Ontology Gas", "6f6e6778", "ong", "Cosmos", "1000000000","ong");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(79, "BNB", "0000000000000000000000000000000000000000", "BNB", "BNB", "1","BNB");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "EBNB", "49fdf74986fbfaa96ae3cfbdf4e3dfda59e8bcec", "BNB", "BNB", "1","BNB");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ontology", "530aae4c0859894023906e28467f2a7f111b6ff3", "ONT Token (ONTX)", "ERC20", "1","ont");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ontology", "0100000000000000000000000000000000000000", "ont", "OEP4", "1","ont");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Ontology", "5a9222225f1bdb135123b74354c7248200c440aa", "ONT_NEP5", "NEP5", "1","ont");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Ontology", "6f6e7478", "ont", "Cosmos", "1","ont");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ethereum", "0000000000000000000000000000000000000000", "ether", "ether", "1000000000000000000","ether");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ethereum", "7009a2f7c8a2e45fa386a6078c7bfeaf518be487", "oETH", "OEP4", "1000000000000000000","ether");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Ethereum", "d7b32de37ad906df80805c2419ff5560d20f9cbf", "ETHxNEO", "NEP5", "1000000000000000000","ether");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Ethereum", "65746878", "ether", "Cosmos", "1000000000000000000","ether");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "erc20", "276788af4a803781267c84692416311de1f761f9", "ERC20 Template (ERC20T)", "ERC20", "1000000000","erc20");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "erc20", "e930755b130dccb25dc3cfee2b2e30d9370c1a75", "ERC20Template", "OEP4", "1000000000","erc20");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "erc20", "657263323078", "erc20", "Cosmos", "1000000000","erc20");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Cosmos", "b25b51d684f7945f7aab43496cb0e87138abdb35", "NEOATOM", "NEP5", "100000000","atom");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Cosmos", "7374616b65", "atom", "Cosmos", "100000000", "atom");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Neo", "7e269f2f33a97c64192e9889faeec72a6fcdb397", "NEO Token (eNEO)", "ERC20", "100000000","neo");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Neo", "13eef3e184d878038317d806796b3af2d9f9b36d", "pNEO", "OEP4", "100000000","neo");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Neo", "c074a05e9dcf0141cbe6b4b3475dd67baf4dcb60", "CNEO", "NEP5", "100000000","neo");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "neo gas", "74f2dc36a68fdc4682034178eb2220729231db76", "CGAS", "NEP5", "100000000","CGAS");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "neo gas", "67617378", "CGAS", "Cosmos", "100000000", "CGAS");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Tether", "ad3f96ae966ad60347f31845b7e4b333104c52fb", "USDT (USDT)", "ERC20", "1000000","usdt");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Tether", "c6f91c11d740d39943b99a6b1c6fd2b5f476e2a3", "oUSDT", "OEP4", "1000000","usdt");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "USD Coin", "0d9c8723b343a8368bebe0b5e89273ff8d712e3c", "USDC", "ERC20", "1000000","usdc");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "USD Coin", "07a12c0a6bdce4df04ef4b2045d1b0fd63a56e25", "pUSDC,", "OEP4", "1000000","usdt");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "557563dc4ed3fd256eba55b9622f53331ab97c2f", "Wrapped BTC (WBTC)", "ERC20", "100000000","wbtc");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Bitcoin", "aede525f05065306423a5522bfcd31b5847ffa52", "oBTC", "OEP4", "100000000","wbtc");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Dai", "8cad2301f7348dfc10c65778197028f432d51e76", "Dai Stablecoin (DAI)", "ERC20", "1000000000000000000","DAI");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Dai", "96cf88356123592835a2fa75068a242260be1791", "pDAI", "OEP4", "1000000000000000000","DAI");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "ONTD", "76130c293aa35bf7b3e5fed1e9ae1e5df12c6a92", "eONTD", "ERC20", "1000000000","ONTD");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "ONTD", "33ae7eae016193ba0fe238b223623bc78faac158", "ONT-decimal", "OEP4", "1000000000","ONTD");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "239100e629a9ca8e0bf45c7892b0fc72d78aa97a", "RENBTC", "ERC20", "100000000","RENBTC");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Bitcoin", "ec547bc4486dea97cb659f1fe73407922f9e63c8", "pRENBTC", "OEP4", "100000000","RENBTC");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Oep4", "3105a14f7956d33a51f12ef3ae50a3f1ef161dff", "EOEP4", "ERC20", "1000000000","EOEP4");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Oep4", "969850e009b5e2a061694f3479ec8e44bc68bcd3", "EOEP4", "OEP4", "1000000000","EOEP4");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "WING", "b60e03e6973b1d0b90a763f5b64c48ca7cb8c2d1", "pWING", "ERC20", "1000000000","WING");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "WING", "ff31ec74d01f7b7d45ed2add930f5d2239f7de33", "WING", "OEP4", "1000000000","WING");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "PAX", "825c227b07153ea450a6607e35c2bf70a9b2fe47", "pPAX", "ERC20", "1000000000","PAX");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "PAX", "5cb420480614fc635f7bfa7f3bd2cd5484004545", "pPAX", "OEP4", "1000000000","PAX");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "SUSD", "21718c0fbd10900565fa57c76e1862cd3f6a4d8e", "SUSD", "ERC20", "1000000000000000000","SUSD");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "SUSD", "37f4497b6f5f511e73843a0bda1042777666f7ec", "pSUSD", "OEP4", "1000000000000000000","SUSD");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "OKB", "776c8db09367615bc741be6e13dec8eabd2bd8bd", "OKB", "ERC20", "1000000000000000000","OKB");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "OKB", "0191f134a3ef0e1eb4f557b6aa0b8bdfd0a5db21", "pOKB", "OEP4", "1000000000000000000","OKB");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Uniswap", "1269d9940a2bfc5ac13c759e7ef1e35fec7278f6", "Uniswap", "ERC20", "1000000000000000000","Uniswap");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Uniswap", "b79d2064947f61070cb68ef26cbc12cbf3b98d9e", "Uniswap", "OEP4", "1000000000000000000","Uniswap");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "HT", "0000000000000000000000000000000000000000", "HT", "HT", "1000000000000000000","HT");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "pHT", "843e9f7a4ba7e062a53d7bbbe85cb35421704616", "HT", "OEP4", "1000000000000000000","HT");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "pNEO", "b119b3b8e5e6eeffbe754b20ee5b8a42809931fb", "NEO", "OEP4", "1000000000000000000","NEO");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "pONTd", "658cabf9c1f71ba0fa64098a7c17e52b94046ece", "ONT", "OEP4", "1000000000","ONT");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pONTd", "01faa7b4157f3a3672892c0ec2773f5400dd2dcf", "ONT", "ERC20", "1000000000","ONT");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "pHT", "930b81bddcd7a793d3541e41d85475a20e169092", "HT", "ERC20", "1000000000000000000","HT");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "pHRC20", "3fdd17abfbb43d29e32746c273085c05d58e2e80", "HRC20", "ERC20", "1000000000000000000","HRC20");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pDAI", "6965731afed27add95be4b0d88ac895ea3eac7ef", "DAI", "ERC20", "1000000000000000000","DAI");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pWING", "b8033859be7553cf9c7657f664fc2243ba5f02ef", "WING", "ERC20", "1000000000","WING");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pUSDC", "63b79d3c4d1f7cbefb80e4b0ab7e8d1545ba10c3", "USDC", "ERC20", "1000000","USDC");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pETH", "0bc413bbde7855ba8114596b6dc53ed8a5c10774", "ETH", "ERC20", "1000000000000000000","ETH");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pWBTC", "b0fcb90b85eab741ba28a132044e93344136ca36", "WBTC", "ERC20", "100000000","WBTC");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pBNB", "33b89f811e8986c5b9d32114a21cc1fd5feb6c78", "BNB", "ERC20", "1000000000000000000","BNB");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pFLM", "b1203bdc2c60521dcbf73dcfc82edbda101ea0a2", "FLM", "ERC20", "1000000000","FLM");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pONTd", "01faa7b4157f3a3672892c0ec2773f5400dd2dcf", "ONTd", "ERC20", "1000000000","ONTd");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "pNEO", "6ef070cb10fc9f66d04a4c387928b268f55b9198", "NEO", "ERC20", "1000000000000000000","NEO");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(7, "HRC20", "faddf0cfb08f92779560db57be6b2c7303aad266", "HRC20", "ERC20", "1000000000000000000","HRC20");




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
  `amount` BIGINT(8) NOT NULL COMMENT '收到的金额',
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
  `amount` BIGINT(8) NOT NULL COMMENT '收到的金额',
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

SET sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''));