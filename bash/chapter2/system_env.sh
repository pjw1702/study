#!/bin/bash

# 예약변수 및 환경변수

# 시스템의 환경정보를 확인할 경우 유용

HOME                # 사용자의 홈 디렉토리
PATH                # 실행 파일을 찾을 디렉토리 경로
FUNCNAME            # 현재 함수 이름
LANG                # 프로그램 사용 시 기본으로 지원되는 언어
PWD                 # 사용자의 현재 작업 중인 디렉토리
TERM                # 로그인 터미널 타입
SHELL               # 로그인해서 사용하는 셸
USER                # 사용자 이름
USERNAME            # 사용자 이름
GROUP               # 사용자 그룹(/etc/passwd 값을 출력)
DISPLAY             # X 디스플레이 이름
COLUMNS             # 현재 터미널이나 윈도우 터미널의 컬럼 수
LINES               # 터미널의 라인 수
PS1                 # 기본 프롬프트 변수
PS2                 # 보조 프롬프트 변수(기본값: >), 명령을 "\"를 사용하여 명령 행 연장 시 사용됨
PS3                 # 셸 스크립트에서 select 사용 시 프롬프트 변수
PS4                 # 셸 스크립트에서 디버깅 모드의 프롬프트 변수(기본값:+)
BASH                # BASH 실행 파일 경로
BASH_VERSION        # 설치된 BASH 버전
BASH_ENV            # 스크립트 실행 시 BASH 시작 파일을 읽을 위치 변수
HISTFILE            # history 파일 경로
HISTFILESIZE        # history 파일 크기
HISTSIZE            # history가 저장되는 갯수
HOSTNAME            # 호스트 이름
HOSTTYPE            # 시스템 하드웨어 종류
MACHTYPE           # 머신 종류(HOSTTYPE과 같은 정보지만 조금더 상세하게 표시됨)
MAIL                # 메일 보관 경로
LOGNAME             # 로그인 이름
TMOUT               # 0이면 제한이 없으며 time 시간 지정 시 지정한 시간 이후 로그아웃
SECONDS             # 스크립트가 실행된 초 단위 시간
UID                 # 사용자 UID
OSTYPE              # 운영체제 종류