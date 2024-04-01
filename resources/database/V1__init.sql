CREATE DATABASE IF NOT EXISTS db_omega;

USE db_omega;

CREATE TABLE IF NOT EXISTS tbl_audit (
     oid integer not null auto_increment,
     source varchar(100),
    source_url varchar(1000),
    resource_url varchar(500),
    system_file_path varchar(1000),
    strategy varchar(20),
    invoker varchar(50),
    channel varchar(100),
    title varchar(255),
    timestamp datetime,
    PRIMARY KEY (oid)
    );