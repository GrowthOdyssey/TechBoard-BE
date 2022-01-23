#!/bin/sh

# テーブル作成スクリプト
psql -U test_user -d tech_board_db -f schema.sql
