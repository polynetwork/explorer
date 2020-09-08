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
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("neo",4,4,4639655,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("cosmos",5,5,0,0,0);

INSERT INTO `chain_contract`(`id`,`contract`) VALUES(0, "0300000000000000000000000000000000000000");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(2, "726532586c50ec9f4080b71f906a3d9779bbd64f");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(3, "0900000000000000000000000000000000000000");
INSERT INTO `chain_contract`(`id`,`contract`) VALUES(4, "80cd0c6fb005da87b78c54dd03c65ef1447195fa");

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




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ontology", "530aae4c0859894023906e28467f2a7f111b6ff3", "ONT Token (ONTX)", "ERC20", "1","ont");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ontology", "0100000000000000000000000000000000000000", "ont", "OEP4", "1","ont");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Ontology", "5a9222225f1bdb135123b74354c7248200c440aa", "ONT_NEP5", "NEP5", "1","ont");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Ontology", "6f6e7478", "ont", "Cosmos", "1","ont");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ethereum", "0000000000000000000000000000000000000000", "ether", "ether", "1000000000000000000","ether");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ethereum", "7448f5f18088c566a3e78c207ea0f06b0aea58b6", "oETH", "OEP4", "1000000000000000000","ether");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Ethereum", "d7b32de37ad906df80805c2419ff5560d20f9cbf", "ETHxNEO", "NEP5", "1000000000000000000","ether");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Ethereum", "65746878", "ether", "Cosmos", "1000000000000000000","ether");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "erc20", "276788af4a803781267c84692416311de1f761f9", "ERC20 Template (ERC20T)", "ERC20", "1000000000","erc20");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "erc20", "e930755b130dccb25dc3cfee2b2e30d9370c1a75", "ERC20Template", "OEP4", "1000000000","erc20");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "erc20", "657263323078", "erc20", "Cosmos", "1000000000","erc20");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Cosmos", "b25b51d684f7945f7aab43496cb0e87138abdb35", "NEOATOM", "NEP5", "100000000","atom");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "Cosmos", "7374616b65", "atom", "Cosmos", "100000000", "atom");



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Neo", "7e269f2f33a97c64192e9889faeec72a6fcdb397", "NEO Token (eNEO)", "ERC20", "100000000","neo");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Neo", "ccee89c6db3e80bd333f89d7d91d4aea1eedae92", "oNEO", "OEP4", "100000000","neo");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "Neo", "c074a05e9dcf0141cbe6b4b3475dd67baf4dcb60", "CNEO", "NEP5", "100000000","neo");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "neo gas", "74f2dc36a68fdc4682034178eb2220729231db76", "CGAS", "NEP5", "100000000","CGAS");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "neo gas", "67617378", "CGAS", "Cosmos", "100000000", "CGAS");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Tether", "ad3f96ae966ad60347f31845b7e4b333104c52fb", "USDT (USDT)", "ERC20", "1000000","usdt");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Tether", "11c60400f54c17df0d8e9c4a38c333b66c1f1c54", "oUSDT", "OEP4", "1000000","usdt");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "557563dc4ed3fd256eBA55B9622f53331ab97c2f", "Wrapped BTC (WBTC)", "ERC20", "100000000","wbtc");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Bitcoin", "c984a8fa8d3b4c9399e36b5e257d9e150d0480a1", "oBTC", "OEP4", "100000000","wbtc");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Dai", "8cad2301f7348dfc10c65778197028f432d51e76", "Dai Stablecoin (DAI)", "ERC20", "1000000000000000000","DAI");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Dai", "3249446364433365f075793e298ee3327c1fcb58", "oDai", "OEP4", "1000000000000000000","DAI");




INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "ONTD", "2e0de81023ea6d32460244f29c57c84ce569e7b7", "ONT-Decimal", "OEP4", "1000000000","ONTD");


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