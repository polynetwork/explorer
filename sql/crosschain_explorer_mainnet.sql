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

INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("poly",0,0,22732,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("btc",1,1,0,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("eth",2,2,10650091,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("ontology",3,3,9300490,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("neo",4,4,6023777,0,0);
INSERT INTO `chain_info`(`xname`,`id`,`xtype`,`height`,`txin`,`txout`) VALUES("switcheo",5,5,202650,0,0);

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



INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ontology Gas", "6a4c89eb9a26a2da34f13f8976daa9fd7526f35c", "eONG", "erc20", "1000000000","ong");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Ontology", "519020fa558a52df57854135345c28024a596b68", "eONT", "erc20", "1000000000","ont");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Bitcoin", "8037dd7161401417d3571b92b86846d34309129a", "pWBTC", "OEP4", "100000000","btc");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ethereum", "a2f89e531e55636d4af1cd044237d2fd5a616c72", "oETH", "OEP4", "1000000000000000000","eth");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Tether", "3e931f60f2cd1387b52f1889dfcaf02a54b2c6a0", "oUSDT", "OEP4", "1000000","usdt");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "USD Coin", "061a07cd393aac289b8ecfda2c3784b637a2fb33", "pUSDC", "OEP4", "1000000","usdc");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Dai", "3f0def1945d7129c5f6625147dcbbaaee402e751", "oDai", "OEP4", "1000000000000000000","dai");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Tether", "dac17f958d2ee523a2206206994597c13d831ec7", "USDT", "erc20", "1000000","usdt");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "2260fac5e5542a773aa44fbcfedf7c193bc2c599", "Wrapped BTC", "erc20", "100000000","btc");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Dai", "6b175474e89094c44da98b954eedeac495271d0f", "DAI", "erc20", "1000000000000000000","dai");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "USD Coin", "a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "USD Coin", "erc20", "1000000","usdc");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Ontology", "6a718e3d8ac0e693225d4b86242de32d56468175", "ontd", "oep4", "1000000000","ont");


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "Bitcoin", "eb4c2781e4eba804ce9a9803c67d0893436bb27d", "renBTC", "ERC20", "100000000","renBTC");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "Bitcoin", "46c3051c553aaeb3724ea69336ec483f39fa91b1", "prenBTC", "OEP4", "100000000","renBTC");


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
