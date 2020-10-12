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


INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(1, "btc", "0000000000000000000000000000000000000011", "btc", "BTC", "100000000","btc");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "btc", "a427c5994c2e00272516140836d21bdc3a538c79", "BTC Token (BTCX)", "ERC20", "100000000","btc");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "btc", "aa0e7aede9c0d5b599b9908237f3df84d7260be5", "btcx", "OEP4", "100000000","btc");
--INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "btc", "3ee29d5cc82771e91383f9ba09c6f5c5878f3f24", "btc", "NEP5", "100000000","btc");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "btc", "62746378", "btc", "Cosmos", "100000000","btc");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "oep4", "27a9206bbf288e469aacc2168e621e503c2bc846", "OEP4 Template (OEP4T)", "ERC20", "1000000000","oep4");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "oep4", "34661f6e90f3d590bbe0281a76935f94bd459532", "OEP4Template", "OEP4", "1000000000","oep4");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "oep4", "6f65703478", "oep4", "Cosmos", "1000000000","oep4");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "ong", "fb37c160cfbd8bd4ba6df6f70e2449b6eb83fc26", "ONG Token (ONGX)", "ERC20", "1000000000","ong");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "ong", "0200000000000000000000000000000000000000", "ong", "OEP4", "1000000000","ong");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "ong", "6f6e6778", "ong", "Cosmos", "1000000000","ong");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "ont", "147cb77cfd0d3a0932d859820d1a3db9686396f9", "ONT Token (ONTX)", "ERC20", "1","ont");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "ont", "0100000000000000000000000000000000000000", "ont", "OEP4", "1","ont");
--INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "ont", "a63d7dffa7718902fda0f64e57f3c5e0c33fd3ff", "ont", "NEP5", "1","ont");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "ont", "6f6e7478", "ont", "Cosmos", "1","ont");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "ether", "0000000000000000000000000000000000000000", "ether", "ether", "1000000000000000000","ether");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "ether", "1fd76f19e621816de34c5c74643ca69aeeab4d42", "ethx", "OEP4", "1000000000000000000","ether");
--INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "ether", "74fac41ad5ad23921a3400e953e1cafb41240d08", "ether", "NEP5", "1000000000000000000","ether");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "ether", "65746878", "ether", "Cosmos", "1000000000000000000","ether");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "erc20", "8142d8936a3571f6f4777716d37b1dd4c1ca1d93", "ERC20 Template (ERC20T)", "ERC20", "1000000000","erc20");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(3, "erc20", "8c8257c1f58533e906f6e24ae06d06be0fc16432", "ERC20Template", "OEP4", "1000000000","erc20");
INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "erc20", "657263323078", "erc20", "Cosmos", "1000000000","erc20");

--INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(2, "neo", "20f307ea523e69d195b3a370fe6496eb50ce281a", "neo", "ERC20", "100000000","neo");
--INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "neo", "c074a05e9dcf0141cbe6b4b3475dd67baf4dcb60", "neo", "NEP5", "100000000","neo");
--INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "neo", "6e656f78", "neo", "Cosmos", "100000000","neo");

--INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(4, "neogas", "74f2dc36a68fdc4682034178eb2220729231db76", "neogas", "NEP5", "100000000","neogas");
--INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "neogas", "67617378", "neogas", "Cosmos", "100000000","neogas");

INSERT INTO `chain_token`(`id`, `xtoken`, `hash`, `xname`, `xtype`,`xprecision`,`xdesc`) VALUES(5, "atom", "7374616b65", "atom", "Cosmos", "100000000", "atom");

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
 INDEX (`height`)
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
  `amount` VARCHAR(32) NOT NULL COMMENT '收到的金额',
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
  `amount` VARCHAR(32) NOT NULL COMMENT '收到的金额',
  PRIMARY KEY (`txhash`),
  INDEX (`asset`, `tt`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `asset_statistic`;
CREATE TABLE `asset_statistic` (
  `xname` VARCHAR(16)  COMMENT '资产名称',
  `addressnum`   INT(4) NOT NULL COMMENT '资产的总地址数',
  `amount`       BIGINT(8)  NOT NULL COMMENT '资产的总价值',
  `amount_btc`   BIGINT(8)  NOT NULL COMMENT '资产的总价值',
  `amount_usd`   BIGINT(8)  NOT NULL COMMENT '资产的总价值',
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

