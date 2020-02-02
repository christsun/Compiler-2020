create database compiler;
use compiler;
create table Dataset_codegen
(
    id             int auto_increment
        primary key,
    cg_uid         varchar(32)               not null,
    cg_sourceCode  text                      not null,
    cg_inputCtx    text                      not null,
    cg_outputCtx   text                      not null,
    cg_outputCode  int                       not null,
    cg_assertion   tinyint(1) default 0      not null,
    cg_timeLimit   float      default -1     not null,
    cg_instLimit   int        default 100000 not null,
    cg_memoryLimit int        default 512    not null
);

create table Dataset_optimize
(
    id                int auto_increment
        primary key,
    optim_uid         varchar(32)               not null,
    optim_sourceCode  text                      not null,
    optim_inputCtx    text                      not null,
    optim_outputCtx   text                      not null,
    optim_outputCode  int                       not null,
    optim_assertion   tinyint(1) default 0      not null,
    optim_timeLimit   float      default -1     not null,
    optim_instLimit   int        default 100000 not null,
    optim_memoryLimit int        default 512    not null
);

create table Dataset_semantic
(
    id               int auto_increment
        primary key,
    sema_uid         varchar(32)               not null,
    sema_sourceCode  text                      not null,
    sema_assertion   tinyint(1) default 0      not null,
    sema_timeLimit   float      default -1     not null,
    sema_instLimit   int        default 100000 not null,
    sema_memoryLimit int        default 512    not null,
    constraint sema_uid_UNIQUE
        unique (sema_uid)
);

create table JudgeDetail
(
    id                int auto_increment
        primary key,
    judge_d_useruuid  varchar(80)            not null,
    judge_d_judger    varchar(16) default '' not null,
    judge_d_judgeTime varchar(20)            not null,
    judge_d_subworkId text                   not null,
    judge_d_testcase  varchar(32)            not null,
    judge_d_result    text                   not null,
    judge_d_type      int         default 0  not null,
    judge_d_githash   varchar(40)            not null
);

create table judgeResult
(
    id                int auto_increment
        primary key,
    judge_p_useruuid  varchar(80)                        not null,
    judge_p_repo      text                               not null,
    judge_p_githash   varchar(40)                        not null,
    judge_p_verdict   int      default 0                 not null,
    judge_p_semantic  text                               null,
    judge_p_codegen   text                               null,
    judge_p_optimize  text                               null,
    judge_p_judgetime datetime default CURRENT_TIMESTAMP null,
    judge_p_judgeid   varchar(80)                        not null
);

create table userDatabase
(
    id           int auto_increment
        primary key,
    stu_uuid     varchar(80) not null,
    stu_id       varchar(12) not null,
    stu_repo     text        not null,
    stu_name     text        not null,
    stu_password varchar(40) not null,
    stu_email    text        not null,
    constraint stu_id_UNIQUE
        unique (stu_id)
);
