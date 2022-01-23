#!/bin/sh

# DB作成スクリプト
psql postgres << EOF
  create user test_user;
  create database tech_board_db owner test_user;
EOF
