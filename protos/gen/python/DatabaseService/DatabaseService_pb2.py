# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: DatabaseService/DatabaseService.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'DatabaseService/DatabaseService.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%DatabaseService/DatabaseService.proto\x12\x07service\"\x07\n\x05\x45mpty\"`\n\x15UpdateUserCardRequest\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x10\n\x08\x66ullName\x18\x02 \x01(\t\x12\x0e\n\x06number\x18\x03 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x04 \x01(\t\x12\x0b\n\x03\x63vv\x18\x05 \x01(\x04\"U\n\x16UpdateUserCardRequest1\x12\x10\n\x08\x66ullName\x18\x01 \x01(\t\x12\x0e\n\x06number\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x03 \x01(\t\x12\x0b\n\x03\x63vv\x18\x04 \x01(\x04\"\x97\x01\n\x16UpdateUserCardResponse\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x10\n\x08\x66ullName\x18\x02 \x01(\t\x12\x0e\n\x06number\x18\x03 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x04 \x01(\t\x12\x0b\n\x03\x63vv\x18\x05 \x01(\x04\x12\x0e\n\x06userId\x18\x06 \x01(\x04\x12\x11\n\tcreatedAt\x18\x07 \x01(\t\x12\x11\n\tupdatedAt\x18\x08 \x01(\t\"\xe5\x01\n\x18\x41\x64\x64\x43\x61rdToCompanyResponse\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05title\x18\x02 \x01(\t\x12\r\n\x05phone\x18\x03 \x01(\t\x12\x0f\n\x07\x61\x64\x64ress\x18\x04 \x01(\t\x12\x0c\n\x04site\x18\x05 \x01(\t\x12\x0b\n\x03inn\x18\x06 \x01(\t\x12\x0b\n\x03kpp\x18\x07 \x01(\t\x12\x0c\n\x04okpo\x18\x08 \x01(\t\x12\"\n\x04\x63\x61rd\x18\t \x01(\x0b\x32\x14.service.CardCompany\x12\x0e\n\x06userId\x18\n \x01(\x04\x12\x11\n\tcreatedAt\x18\x0b \x01(\t\x12\x11\n\tupdatedAt\x18\x0c \x01(\t\"J\n\x17\x41\x64\x64\x43\x61rdToCompanyRequest\x12/\n\x04\x63\x61rd\x18\x01 \x01(\x0b\x32!.service.CreateCardCompanyRequest\".\n\x1c\x44\x65leteSessionByUserIdRequest\x12\x0e\n\x06userId\x18\x01 \x01(\x04\"&\n\x18\x44\x65leteSessionByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"u\n\x1b\x44\x65leteSessionByModelRequest\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x0e\n\x06userId\x18\x02 \x01(\x04\x12\x14\n\x0crefreshToken\x18\x03 \x01(\t\x12\x11\n\tcreatedAt\x18\x04 \x01(\t\x12\x11\n\tupdatedAt\x18\x05 \x01(\t\"I\n!ChangeRefreshTokenByUserIdRequest\x12\x0e\n\x06userId\x18\x01 \x01(\x04\x12\x14\n\x0crefreshToken\x18\x02 \x01(\t\"7\n\"ChangeRefreshTokenByUserIdResponse\x12\x11\n\taccessory\x18\x01 \x01(\x08\"A\n\x1d\x43hangeRefreshTokenByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x14\n\x0crefreshToken\x18\x02 \x01(\t\"3\n\x1e\x43hangeRefreshTokenByIdResponse\x12\x11\n\taccessory\x18\x01 \x01(\x08\"-\n\x1b\x46indSessionsByUserIdRequest\x12\x0e\n\x06userId\x18\x01 \x01(\x04\"v\n\x1c\x46indSessionsByUserIdResponse\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x0e\n\x06userId\x18\x02 \x01(\x04\x12\x14\n\x0crefreshToken\x18\x03 \x01(\t\x12\x11\n\tcreatedAt\x18\x04 \x01(\t\x12\x11\n\tupdatedAt\x18\x05 \x01(\t\"%\n\x17\x46indSessionsByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"r\n\x18\x46indSessionsByIdResponse\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x0e\n\x06userId\x18\x02 \x01(\x04\x12\x14\n\x0crefreshToken\x18\x03 \x01(\t\x12\x11\n\tcreatedAt\x18\x04 \x01(\t\x12\x11\n\tupdatedAt\x18\x05 \x01(\t\"D\n\x10SessionsResponse\x12\x30\n\x08sessions\x18\x01 \x03(\x0b\x32\x1e.service.CreateSessionResponse\"<\n\x14\x43reateSessionRequest\x12\x0e\n\x06userId\x18\x01 \x01(\x04\x12\x14\n\x0crefreshToken\x18\x02 \x01(\t\"o\n\x15\x43reateSessionResponse\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x0e\n\x06userId\x18\x02 \x01(\x04\x12\x14\n\x0crefreshToken\x18\x03 \x01(\t\x12\x11\n\tcreatedAt\x18\x04 \x01(\t\x12\x11\n\tupdatedAt\x18\x05 \x01(\t\"$\n\x16\x46indUserCompanyRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"\x19\n\tHTTPCodes\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x03\"@\n\x14\x41\x64\x64\x43\x61rdToUserRequest\x12(\n\x04\x63\x61rd\x18\x01 \x01(\x0b\x32\x1a.service.CreateCardRequest\"\x97\x01\n\x16UpdateDonationsRequest\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05title\x18\x02 \x01(\t\x12\x0e\n\x06\x61mount\x18\x03 \x01(\x02\x12\x1c\n\x05wards\x18\x04 \x01(\x0b\x32\r.service.Ward\x12\x0e\n\x06userId\x18\x05 \x01(\x04\x12\x11\n\tcreatedAt\x18\x06 \x01(\t\x12\x11\n\tupdatedAt\x18\x07 \x01(\t\"\x9d\x01\n\x1c\x44\x65leteDonationByModelRequest\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05title\x18\x02 \x01(\t\x12\x0e\n\x06\x61mount\x18\x03 \x01(\x02\x12\x1c\n\x05wards\x18\x04 \x01(\x0b\x32\r.service.Ward\x12\x0e\n\x06userId\x18\x05 \x01(\x04\x12\x11\n\tcreatedAt\x18\x06 \x01(\t\x12\x11\n\tupdatedAt\x18\x07 \x01(\t\"\xa2\x02\n\x15\x41\x64\x64\x43\x61rdToUserResponse\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x10\n\x08username\x18\x03 \x01(\t\x12\x10\n\x08password\x18\x04 \x01(\t\x12\r\n\x05phone\x18\x05 \x01(\t\x12\x1b\n\x04\x63\x61rd\x18\x06 \x03(\x0b\x32\r.service.Card\x12\x0c\n\x04role\x18\x07 \x01(\t\x12!\n\x07\x63ompany\x18\x08 \x01(\x0b\x32\x10.service.Company\x12\x0c\n\x04type\x18\t \x01(\x04\x12%\n\tdonations\x18\n \x01(\x0b\x32\x12.service.Donations\x12\x12\n\nAvatarPath\x18\x0b \x01(\t\x12\x11\n\tcreatedAt\x18\x0c \x01(\t\x12\x11\n\tupdatedAt\x18\r \x01(\t\"\x98\x01\n\x17\x43reateDonationsResponse\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05title\x18\x02 \x01(\t\x12\x0e\n\x06\x61mount\x18\x03 \x01(\x02\x12\x1c\n\x05wards\x18\x04 \x01(\x0b\x32\r.service.Ward\x12\x0e\n\x06userId\x18\x05 \x01(\x04\x12\x11\n\tcreatedAt\x18\x06 \x01(\t\x12\x11\n\tupdatedAt\x18\x07 \x01(\t\"#\n\x15\x44\x65leteWardByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"!\n\x13\x46indWardByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"i\n\x11\x43reateWardRequest\x12\r\n\x05title\x18\x02 \x01(\t\x12\x10\n\x08\x66ullName\x18\x03 \x01(\t\x12\x0c\n\x04want\x18\x04 \x01(\t\x12\x11\n\tnecessary\x18\x05 \x01(\x02\x12\x12\n\ndonationId\x18\x06 \x01(\x04\"-\n\rWardsResponse\x12\x1c\n\x05wards\x18\x01 \x03(\x0b\x32\r.service.Ward\"\'\n\x19\x44\x65leteDonationByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"%\n\x17\x46indDonationByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"9\n\x19\x46indDonationWardsResponse\x12\x1c\n\x05wards\x18\x01 \x01(\x0b\x32\r.service.Ward\"&\n\x18\x46indDonationWardsRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"e\n\x16\x43reateDonationsRequest\x12\r\n\x05title\x18\x01 \x01(\t\x12\x0e\n\x06\x61mount\x18\x02 \x01(\x02\x12\x1c\n\x05wards\x18\x03 \x01(\x0b\x32\r.service.Ward\x12\x0e\n\x06userId\x18\x04 \x01(\x04\":\n\x11\x44onationsResponse\x12%\n\tdonations\x18\x01 \x03(\x0b\x32\x12.service.Donations\"*\n\x1c\x44\x65leteCardCompanyByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"(\n\x1a\x46indCardCompanyByIDRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"j\n\x18\x43reateCardCompanyRequest\x12\x10\n\x08\x66ullName\x18\x02 \x01(\t\x12\x0e\n\x06number\x18\x03 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x04 \x01(\t\x12\x0b\n\x03\x63vv\x18\x05 \x01(\x04\x12\x11\n\tcompanyId\x18\x06 \x01(\x04\"=\n\x16\x43\x61rdsCompaniesResponse\x12#\n\x05\x63\x61rds\x18\x01 \x03(\x0b\x32\x14.service.CardCompany\"#\n\x15\x44\x65leteCardByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"!\n\x13\x46indCardByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"`\n\x11\x43reateCardRequest\x12\x10\n\x08\x66ullName\x18\x01 \x01(\t\x12\x0e\n\x06number\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x03 \x01(\t\x12\x0b\n\x03\x63vv\x18\x04 \x01(\x04\x12\x0e\n\x06userId\x18\x05 \x01(\x04\"-\n\rCardsResponse\x12\x1c\n\x05\x63\x61rds\x18\x01 \x03(\x0b\x32\r.service.Card\"9\n\x14UpdateCompanyRequest\x12!\n\x07\x63ompany\x18\x01 \x01(\x0b\x32\x10.service.Company\"@\n\x1b\x44\x65leteCompanyByModelRequest\x12!\n\x07\x63ompany\x18\x01 \x01(\x0b\x32\x10.service.Company\"&\n\x18\x44\x65leteCompanyByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"$\n\x16\x46indCompanyCardRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"$\n\x16\x46indCompanyByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\",\n\x1b\x46indCompanyByIdPhoneRequest\x12\r\n\x05phone\x18\x01 \x01(\t\"\xaf\x01\n\x14\x43reateCompanyRequest\x12\r\n\x05title\x18\x01 \x01(\t\x12\r\n\x05phone\x18\x02 \x01(\t\x12\x0f\n\x07\x61\x64\x64ress\x18\x03 \x01(\t\x12\x0c\n\x04site\x18\x04 \x01(\t\x12\x0b\n\x03inn\x18\x05 \x01(\t\x12\x0b\n\x03kpp\x18\x06 \x01(\t\x12\x0c\n\x04okpo\x18\x07 \x01(\t\x12\"\n\x04\x63\x61rd\x18\x08 \x01(\x0b\x32\x14.service.CardCompany\x12\x0e\n\x06userId\x18\t \x01(\x04\"8\n\x11\x43ompaniesResponse\x12#\n\tcompanies\x18\x01 \x03(\x0b\x32\x10.service.Company\"\x9e\x02\n\x11UpdateUserRequest\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x10\n\x08username\x18\x03 \x01(\t\x12\x10\n\x08password\x18\x04 \x01(\t\x12\r\n\x05phone\x18\x05 \x01(\t\x12\x1b\n\x04\x63\x61rd\x18\x06 \x03(\x0b\x32\r.service.Card\x12\x0c\n\x04role\x18\x07 \x01(\t\x12!\n\x07\x63ompany\x18\x08 \x01(\x0b\x32\x10.service.Company\x12\x0c\n\x04type\x18\t \x01(\x04\x12%\n\tdonations\x18\n \x03(\x0b\x32\x12.service.Donations\x12\x12\n\nAvatarPath\x18\x0b \x01(\t\x12\x11\n\tcreatedAt\x18\x0c \x01(\t\x12\x11\n\tupdatedAt\x18\r \x01(\t\"E\n\x18\x44\x65leteUserByModelRequest\x12)\n\x04user\x18\x01 \x01(\x0b\x32\x1b.service.CreateUserResponse\"#\n\x15\x44\x65leteUserByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"!\n\x13\x46indUserCardRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"4\n\x14\x46indUserCardResponse\x12\x1c\n\x05\x63\x61rds\x18\x01 \x03(\x0b\x32\r.service.Card\"1\n\x15\x43hangeUserTypeRequest\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x0c\n\x04type\x18\x02 \x01(\x04\"+\n\x16\x43hangeUserTypeResponse\x12\x11\n\taccessory\x18\x01 \x01(\x08\"\'\n\x16\x46indUserByPhoneRequest\x12\r\n\x05phone\x18\x01 \x01(\t\"!\n\x13\x46indUserByIdRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"&\n\x18\x46indUserDonationsRequest\x12\n\n\x02id\x18\x01 \x01(\x04\"B\n\x19\x46indUserDonationsResponse\x12%\n\tdonations\x18\x01 \x03(\x0b\x32\x12.service.Donations\"\'\n\x16\x46indUserByEmailRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\";\n\rUsersResponse\x12*\n\x05users\x18\x01 \x03(\x0b\x32\x1b.service.CreateUserResponse\"$\n\x13UserIsExistsRequest\x12\r\n\x05phone\x18\x01 \x01(\t\"(\n\x14UserIsExistsResponse\x12\x10\n\x08isExists\x18\x01 \x01(\x08\"9\n\x16\x43omparePasswordRequest\x12\r\n\x05phone\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\",\n\x17\x43omparePasswordResponse\x12\x11\n\taccessory\x18\x01 \x01(\x08\"\xd4\x01\n\x07\x43ompany\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05title\x18\x02 \x01(\t\x12\r\n\x05phone\x18\x03 \x01(\t\x12\x0f\n\x07\x61\x64\x64ress\x18\x04 \x01(\t\x12\x0c\n\x04site\x18\x05 \x01(\t\x12\x0b\n\x03inn\x18\x06 \x01(\t\x12\x0b\n\x03kpp\x18\x07 \x01(\t\x12\x0c\n\x04okpo\x18\x08 \x01(\t\x12\"\n\x04\x63\x61rd\x18\t \x01(\x0b\x32\x14.service.CardCompany\x12\x0e\n\x06userId\x18\n \x01(\x04\x12\x11\n\tcreatedAt\x18\x0b \x01(\t\x12\x11\n\tupdatedAt\x18\x0c \x01(\t\"\x8f\x01\n\x0b\x43\x61rdCompany\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x10\n\x08\x66ullName\x18\x02 \x01(\t\x12\x0e\n\x06number\x18\x03 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x04 \x01(\t\x12\x0b\n\x03\x63vv\x18\x05 \x01(\x04\x12\x11\n\tcompanyId\x18\x06 \x01(\x04\x12\x11\n\tcreatedAt\x18\x07 \x01(\t\x12\x11\n\tupdatedAt\x18\x08 \x01(\t\"\x85\x01\n\x04\x43\x61rd\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x10\n\x08\x66ullName\x18\x02 \x01(\t\x12\x0e\n\x06number\x18\x03 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x04 \x01(\t\x12\x0b\n\x03\x63vv\x18\x05 \x01(\x04\x12\x0e\n\x06userId\x18\x06 \x01(\x04\x12\x11\n\tcreatedAt\x18\x07 \x01(\t\x12\x11\n\tupdatedAt\x18\x08 \x01(\t\"\xa2\x01\n\x04Ward\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05title\x18\x02 \x01(\t\x12\x10\n\x08\x66ullName\x18\x03 \x01(\t\x12\x0c\n\x04want\x18\x04 \x01(\t\x12\x11\n\tnecessary\x18\x05 \x01(\x02\x12\x12\n\ndonationId\x18\x06 \x01(\x04\x12\x12\n\nAvatarPath\x18\x07 \x01(\t\x12\x11\n\tcreatedAt\x18\x08 \x01(\t\x12\x11\n\tupdatedAt\x18\t \x01(\t\"\x89\x01\n\tDonations\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05title\x18\x02 \x01(\t\x12\x0e\n\x06\x61mount\x18\x03 \x01(\x02\x12\x1b\n\x04ward\x18\x04 \x01(\x0b\x32\r.service.Ward\x12\x0e\n\x06userId\x18\x05 \x01(\x04\x12\x11\n\tcreatedAt\x18\x06 \x01(\t\x12\x11\n\tupdatedAt\x18\x07 \x01(\t\"\xff\x01\n\x11\x43reateUserRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\x10\n\x08username\x18\x02 \x01(\t\x12\x10\n\x08password\x18\x03 \x01(\t\x12\r\n\x05phone\x18\x04 \x01(\t\x12(\n\x04\x63\x61rd\x18\x05 \x03(\x0b\x32\x1a.service.CreateCardRequest\x12\x0c\n\x04role\x18\x06 \x01(\t\x12.\n\x07\x63ompany\x18\x07 \x01(\x0b\x32\x1d.service.CreateCompanyRequest\x12\x0c\n\x04type\x18\x08 \x01(\x04\x12\x32\n\tdonations\x18\t \x03(\x0b\x32\x1f.service.CreateDonationsRequest\"\x9f\x02\n\x12\x43reateUserResponse\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x10\n\x08username\x18\x03 \x01(\t\x12\x10\n\x08password\x18\x04 \x01(\t\x12\r\n\x05phone\x18\x05 \x01(\t\x12\x1b\n\x04\x63\x61rd\x18\x06 \x03(\x0b\x32\r.service.Card\x12\x0c\n\x04role\x18\x07 \x01(\t\x12!\n\x07\x63ompany\x18\x08 \x01(\x0b\x32\x10.service.Company\x12\x0c\n\x04type\x18\t \x01(\x04\x12%\n\tdonations\x18\n \x03(\x0b\x32\x12.service.Donations\x12\x12\n\nAvatarPath\x18\x0b \x01(\t\x12\x11\n\tcreatedAt\x18\x0c \x01(\t\x12\x11\n\tupdatedAt\x18\r \x01(\t\")\n\rIsRoleRequest\x12\n\n\x02id\x18\x01 \x01(\x04\x12\x0c\n\x04role\x18\x02 \x01(\t\"#\n\x0eIsRoleResponse\x12\x11\n\taccessory\x18\x01 \x01(\x08\x32\xaa\"\n\x0f\x44\x61tabaseService\x12\x45\n\nCreateUser\x12\x1a.service.CreateUserRequest\x1a\x1b.service.CreateUserResponse\x12/\n\x05Users\x12\x0e.service.Empty\x1a\x16.service.UsersResponse\x12\x39\n\x06IsRole\x12\x16.service.IsRoleRequest\x1a\x17.service.IsRoleResponse\x12T\n\x0f\x43omparePassword\x12\x1f.service.ComparePasswordRequest\x1a .service.ComparePasswordResponse\x12K\n\x0cUserIsExists\x12\x1c.service.UserIsExistsRequest\x1a\x1d.service.UserIsExistsResponse\x12I\n\x0c\x46indUserById\x12\x1c.service.FindUserByIdRequest\x1a\x1b.service.CreateUserResponse\x12O\n\x0f\x46indUserByEmail\x12\x1f.service.FindUserByEmailRequest\x1a\x1b.service.CreateUserResponse\x12O\n\x0f\x46indUserByPhone\x12\x1f.service.FindUserByPhoneRequest\x1a\x1b.service.CreateUserResponse\x12Q\n\x0e\x43hangeUserType\x12\x1e.service.ChangeUserTypeRequest\x1a\x1f.service.ChangeUserTypeResponse\x12\x44\n\x0f\x46indUserCompany\x12\x1f.service.FindUserCompanyRequest\x1a\x10.service.Company\x12Z\n\x11\x46indUserDonations\x12!.service.FindUserDonationsRequest\x1a\".service.FindUserDonationsResponse\x12K\n\x0c\x46indUserCard\x12\x1c.service.FindUserCardRequest\x1a\x1d.service.FindUserCardResponse\x12N\n\rAddCardToUser\x12\x1d.service.AddCardToUserRequest\x1a\x1e.service.AddCardToUserResponse\x12J\n\x11\x44\x65leteUserByModel\x12!.service.DeleteUserByModelRequest\x1a\x12.service.HTTPCodes\x12\x44\n\x0e\x44\x65leteUserById\x12\x1e.service.DeleteUserByIdRequest\x1a\x12.service.HTTPCodes\x12\x45\n\nUpdateUser\x12\x1a.service.UpdateUserRequest\x1a\x1b.service.CreateUserResponse\x12\x37\n\tCompanies\x12\x0e.service.Empty\x1a\x1a.service.CompaniesResponse\x12@\n\rCreateCompany\x12\x1d.service.CreateCompanyRequest\x1a\x10.service.Company\x12\x44\n\x0f\x46indCompanyById\x12\x1f.service.FindCompanyByIdRequest\x1a\x10.service.Company\x12L\n\x12\x46indCompanyByPhone\x12$.service.FindCompanyByIdPhoneRequest\x1a\x10.service.Company\x12H\n\x0f\x46indCompanyCard\x12\x1f.service.FindCompanyCardRequest\x1a\x14.service.CardCompany\x12P\n\x14\x44\x65leteCompanyByModel\x12$.service.DeleteCompanyByModelRequest\x1a\x12.service.HTTPCodes\x12J\n\x11\x44\x65leteCompanyById\x12!.service.DeleteCompanyByIdRequest\x1a\x12.service.HTTPCodes\x12\x42\n\rUpdateCompany\x12\x1d.service.UpdateCompanyRequest\x1a\x12.service.HTTPCodes\x12W\n\x10\x41\x64\x64\x43\x61rdToCompany\x12 .service.AddCardToCompanyRequest\x1a!.service.AddCardToCompanyResponse\x12/\n\x05\x43\x61rds\x12\x0e.service.Empty\x1a\x16.service.CardsResponse\x12\x37\n\nCreateCard\x12\x1a.service.CreateCardRequest\x1a\r.service.Card\x12;\n\x0c\x46indCardById\x12\x1c.service.FindCardByIdRequest\x1a\r.service.Card\x12\x36\n\x11\x44\x65leteCardByModel\x12\r.service.Card\x1a\x12.service.HTTPCodes\x12\x44\n\x0e\x44\x65leteCardById\x12\x1e.service.DeleteCardByIdRequest\x1a\x12.service.HTTPCodes\x12M\n\nUpdateCard\x12\x1e.service.UpdateUserCardRequest\x1a\x1f.service.UpdateUserCardResponse\x12\x41\n\x0e\x43\x61rdsCompanies\x12\x0e.service.Empty\x1a\x1f.service.CardsCompaniesResponse\x12L\n\x11\x43reateCardCompany\x12!.service.CreateCardCompanyRequest\x1a\x14.service.CardCompany\x12P\n\x13\x46indCardCompanyByID\x12#.service.FindCardCompanyByIDRequest\x1a\x14.service.CardCompany\x12\x44\n\x18\x44\x65leteCardCompanyByModel\x12\x14.service.CardCompany\x1a\x12.service.HTTPCodes\x12R\n\x15\x44\x65leteCardCompanyById\x12%.service.DeleteCardCompanyByIdRequest\x1a\x12.service.HTTPCodes\x12?\n\x11UpdateCardCompany\x12\x14.service.CardCompany\x1a\x14.service.CardCompany\x12\x37\n\tDonations\x12\x0e.service.Empty\x1a\x1a.service.DonationsResponse\x12T\n\x0f\x43reateDonations\x12\x1f.service.CreateDonationsRequest\x1a .service.CreateDonationsResponse\x12Z\n\x11\x46indDonationWards\x12!.service.FindDonationWardsRequest\x1a\".service.FindDonationWardsResponse\x12V\n\x10\x46indDonationById\x12 .service.FindDonationByIdRequest\x1a .service.CreateDonationsResponse\x12R\n\x15\x44\x65leteDonationByModel\x12%.service.DeleteDonationByModelRequest\x1a\x12.service.HTTPCodes\x12L\n\x12\x44\x65leteDonationById\x12\".service.DeleteDonationByIdRequest\x1a\x12.service.HTTPCodes\x12S\n\x0eUpdateDonation\x12\x1f.service.UpdateDonationsRequest\x1a .service.CreateDonationsResponse\x12/\n\x05Wards\x12\x0e.service.Empty\x1a\x16.service.WardsResponse\x12\x37\n\nCreateWard\x12\x1a.service.CreateWardRequest\x1a\r.service.Ward\x12;\n\x0c\x46indWardById\x12\x1c.service.FindWardByIdRequest\x1a\r.service.Ward\x12\x36\n\x11\x44\x65leteWardByModel\x12\r.service.Ward\x1a\x12.service.HTTPCodes\x12\x44\n\x0e\x44\x65leteWardById\x12\x1e.service.DeleteWardByIdRequest\x1a\x12.service.HTTPCodes\x12*\n\nUpdateWard\x12\r.service.Ward\x1a\r.service.Ward\x12O\n\x0e\x43reateSessions\x12\x1d.service.CreateSessionRequest\x1a\x1e.service.CreateSessionResponse\x12\x35\n\x08Sessions\x12\x0e.service.Empty\x1a\x19.service.SessionsResponse\x12W\n\x10\x46indSessionsById\x12 .service.FindSessionsByIdRequest\x1a!.service.FindSessionsByIdResponse\x12\x63\n\x14\x46indSessionsByUserId\x12$.service.FindSessionsByUserIdRequest\x1a%.service.FindSessionsByUserIdResponse\x12i\n\x16\x43hangeRefreshTokenById\x12&.service.ChangeRefreshTokenByIdRequest\x1a\'.service.ChangeRefreshTokenByIdResponse\x12u\n\x1a\x43hangeRefreshTokenByUserId\x12*.service.ChangeRefreshTokenByUserIdRequest\x1a+.service.ChangeRefreshTokenByUserIdResponse\x12P\n\x14\x44\x65leteSessionByModel\x12$.service.DeleteSessionByModelRequest\x1a\x12.service.HTTPCodes\x12J\n\x11\x44\x65leteSessionById\x12!.service.DeleteSessionByIdRequest\x1a\x12.service.HTTPCodes\x12R\n\x15\x44\x65leteSessionByUserId\x12%.service.DeleteSessionByUserIdRequest\x1a\x12.service.HTTPCodesB/Z-collapse.DatabaseService.v1;DatabaseServicev1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'DatabaseService.DatabaseService_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z-collapse.DatabaseService.v1;DatabaseServicev1'
  _globals['_EMPTY']._serialized_start=50
  _globals['_EMPTY']._serialized_end=57
  _globals['_UPDATEUSERCARDREQUEST']._serialized_start=59
  _globals['_UPDATEUSERCARDREQUEST']._serialized_end=155
  _globals['_UPDATEUSERCARDREQUEST1']._serialized_start=157
  _globals['_UPDATEUSERCARDREQUEST1']._serialized_end=242
  _globals['_UPDATEUSERCARDRESPONSE']._serialized_start=245
  _globals['_UPDATEUSERCARDRESPONSE']._serialized_end=396
  _globals['_ADDCARDTOCOMPANYRESPONSE']._serialized_start=399
  _globals['_ADDCARDTOCOMPANYRESPONSE']._serialized_end=628
  _globals['_ADDCARDTOCOMPANYREQUEST']._serialized_start=630
  _globals['_ADDCARDTOCOMPANYREQUEST']._serialized_end=704
  _globals['_DELETESESSIONBYUSERIDREQUEST']._serialized_start=706
  _globals['_DELETESESSIONBYUSERIDREQUEST']._serialized_end=752
  _globals['_DELETESESSIONBYIDREQUEST']._serialized_start=754
  _globals['_DELETESESSIONBYIDREQUEST']._serialized_end=792
  _globals['_DELETESESSIONBYMODELREQUEST']._serialized_start=794
  _globals['_DELETESESSIONBYMODELREQUEST']._serialized_end=911
  _globals['_CHANGEREFRESHTOKENBYUSERIDREQUEST']._serialized_start=913
  _globals['_CHANGEREFRESHTOKENBYUSERIDREQUEST']._serialized_end=986
  _globals['_CHANGEREFRESHTOKENBYUSERIDRESPONSE']._serialized_start=988
  _globals['_CHANGEREFRESHTOKENBYUSERIDRESPONSE']._serialized_end=1043
  _globals['_CHANGEREFRESHTOKENBYIDREQUEST']._serialized_start=1045
  _globals['_CHANGEREFRESHTOKENBYIDREQUEST']._serialized_end=1110
  _globals['_CHANGEREFRESHTOKENBYIDRESPONSE']._serialized_start=1112
  _globals['_CHANGEREFRESHTOKENBYIDRESPONSE']._serialized_end=1163
  _globals['_FINDSESSIONSBYUSERIDREQUEST']._serialized_start=1165
  _globals['_FINDSESSIONSBYUSERIDREQUEST']._serialized_end=1210
  _globals['_FINDSESSIONSBYUSERIDRESPONSE']._serialized_start=1212
  _globals['_FINDSESSIONSBYUSERIDRESPONSE']._serialized_end=1330
  _globals['_FINDSESSIONSBYIDREQUEST']._serialized_start=1332
  _globals['_FINDSESSIONSBYIDREQUEST']._serialized_end=1369
  _globals['_FINDSESSIONSBYIDRESPONSE']._serialized_start=1371
  _globals['_FINDSESSIONSBYIDRESPONSE']._serialized_end=1485
  _globals['_SESSIONSRESPONSE']._serialized_start=1487
  _globals['_SESSIONSRESPONSE']._serialized_end=1555
  _globals['_CREATESESSIONREQUEST']._serialized_start=1557
  _globals['_CREATESESSIONREQUEST']._serialized_end=1617
  _globals['_CREATESESSIONRESPONSE']._serialized_start=1619
  _globals['_CREATESESSIONRESPONSE']._serialized_end=1730
  _globals['_FINDUSERCOMPANYREQUEST']._serialized_start=1732
  _globals['_FINDUSERCOMPANYREQUEST']._serialized_end=1768
  _globals['_HTTPCODES']._serialized_start=1770
  _globals['_HTTPCODES']._serialized_end=1795
  _globals['_ADDCARDTOUSERREQUEST']._serialized_start=1797
  _globals['_ADDCARDTOUSERREQUEST']._serialized_end=1861
  _globals['_UPDATEDONATIONSREQUEST']._serialized_start=1864
  _globals['_UPDATEDONATIONSREQUEST']._serialized_end=2015
  _globals['_DELETEDONATIONBYMODELREQUEST']._serialized_start=2018
  _globals['_DELETEDONATIONBYMODELREQUEST']._serialized_end=2175
  _globals['_ADDCARDTOUSERRESPONSE']._serialized_start=2178
  _globals['_ADDCARDTOUSERRESPONSE']._serialized_end=2468
  _globals['_CREATEDONATIONSRESPONSE']._serialized_start=2471
  _globals['_CREATEDONATIONSRESPONSE']._serialized_end=2623
  _globals['_DELETEWARDBYIDREQUEST']._serialized_start=2625
  _globals['_DELETEWARDBYIDREQUEST']._serialized_end=2660
  _globals['_FINDWARDBYIDREQUEST']._serialized_start=2662
  _globals['_FINDWARDBYIDREQUEST']._serialized_end=2695
  _globals['_CREATEWARDREQUEST']._serialized_start=2697
  _globals['_CREATEWARDREQUEST']._serialized_end=2802
  _globals['_WARDSRESPONSE']._serialized_start=2804
  _globals['_WARDSRESPONSE']._serialized_end=2849
  _globals['_DELETEDONATIONBYIDREQUEST']._serialized_start=2851
  _globals['_DELETEDONATIONBYIDREQUEST']._serialized_end=2890
  _globals['_FINDDONATIONBYIDREQUEST']._serialized_start=2892
  _globals['_FINDDONATIONBYIDREQUEST']._serialized_end=2929
  _globals['_FINDDONATIONWARDSRESPONSE']._serialized_start=2931
  _globals['_FINDDONATIONWARDSRESPONSE']._serialized_end=2988
  _globals['_FINDDONATIONWARDSREQUEST']._serialized_start=2990
  _globals['_FINDDONATIONWARDSREQUEST']._serialized_end=3028
  _globals['_CREATEDONATIONSREQUEST']._serialized_start=3030
  _globals['_CREATEDONATIONSREQUEST']._serialized_end=3131
  _globals['_DONATIONSRESPONSE']._serialized_start=3133
  _globals['_DONATIONSRESPONSE']._serialized_end=3191
  _globals['_DELETECARDCOMPANYBYIDREQUEST']._serialized_start=3193
  _globals['_DELETECARDCOMPANYBYIDREQUEST']._serialized_end=3235
  _globals['_FINDCARDCOMPANYBYIDREQUEST']._serialized_start=3237
  _globals['_FINDCARDCOMPANYBYIDREQUEST']._serialized_end=3277
  _globals['_CREATECARDCOMPANYREQUEST']._serialized_start=3279
  _globals['_CREATECARDCOMPANYREQUEST']._serialized_end=3385
  _globals['_CARDSCOMPANIESRESPONSE']._serialized_start=3387
  _globals['_CARDSCOMPANIESRESPONSE']._serialized_end=3448
  _globals['_DELETECARDBYIDREQUEST']._serialized_start=3450
  _globals['_DELETECARDBYIDREQUEST']._serialized_end=3485
  _globals['_FINDCARDBYIDREQUEST']._serialized_start=3487
  _globals['_FINDCARDBYIDREQUEST']._serialized_end=3520
  _globals['_CREATECARDREQUEST']._serialized_start=3522
  _globals['_CREATECARDREQUEST']._serialized_end=3618
  _globals['_CARDSRESPONSE']._serialized_start=3620
  _globals['_CARDSRESPONSE']._serialized_end=3665
  _globals['_UPDATECOMPANYREQUEST']._serialized_start=3667
  _globals['_UPDATECOMPANYREQUEST']._serialized_end=3724
  _globals['_DELETECOMPANYBYMODELREQUEST']._serialized_start=3726
  _globals['_DELETECOMPANYBYMODELREQUEST']._serialized_end=3790
  _globals['_DELETECOMPANYBYIDREQUEST']._serialized_start=3792
  _globals['_DELETECOMPANYBYIDREQUEST']._serialized_end=3830
  _globals['_FINDCOMPANYCARDREQUEST']._serialized_start=3832
  _globals['_FINDCOMPANYCARDREQUEST']._serialized_end=3868
  _globals['_FINDCOMPANYBYIDREQUEST']._serialized_start=3870
  _globals['_FINDCOMPANYBYIDREQUEST']._serialized_end=3906
  _globals['_FINDCOMPANYBYIDPHONEREQUEST']._serialized_start=3908
  _globals['_FINDCOMPANYBYIDPHONEREQUEST']._serialized_end=3952
  _globals['_CREATECOMPANYREQUEST']._serialized_start=3955
  _globals['_CREATECOMPANYREQUEST']._serialized_end=4130
  _globals['_COMPANIESRESPONSE']._serialized_start=4132
  _globals['_COMPANIESRESPONSE']._serialized_end=4188
  _globals['_UPDATEUSERREQUEST']._serialized_start=4191
  _globals['_UPDATEUSERREQUEST']._serialized_end=4477
  _globals['_DELETEUSERBYMODELREQUEST']._serialized_start=4479
  _globals['_DELETEUSERBYMODELREQUEST']._serialized_end=4548
  _globals['_DELETEUSERBYIDREQUEST']._serialized_start=4550
  _globals['_DELETEUSERBYIDREQUEST']._serialized_end=4585
  _globals['_FINDUSERCARDREQUEST']._serialized_start=4587
  _globals['_FINDUSERCARDREQUEST']._serialized_end=4620
  _globals['_FINDUSERCARDRESPONSE']._serialized_start=4622
  _globals['_FINDUSERCARDRESPONSE']._serialized_end=4674
  _globals['_CHANGEUSERTYPEREQUEST']._serialized_start=4676
  _globals['_CHANGEUSERTYPEREQUEST']._serialized_end=4725
  _globals['_CHANGEUSERTYPERESPONSE']._serialized_start=4727
  _globals['_CHANGEUSERTYPERESPONSE']._serialized_end=4770
  _globals['_FINDUSERBYPHONEREQUEST']._serialized_start=4772
  _globals['_FINDUSERBYPHONEREQUEST']._serialized_end=4811
  _globals['_FINDUSERBYIDREQUEST']._serialized_start=4813
  _globals['_FINDUSERBYIDREQUEST']._serialized_end=4846
  _globals['_FINDUSERDONATIONSREQUEST']._serialized_start=4848
  _globals['_FINDUSERDONATIONSREQUEST']._serialized_end=4886
  _globals['_FINDUSERDONATIONSRESPONSE']._serialized_start=4888
  _globals['_FINDUSERDONATIONSRESPONSE']._serialized_end=4954
  _globals['_FINDUSERBYEMAILREQUEST']._serialized_start=4956
  _globals['_FINDUSERBYEMAILREQUEST']._serialized_end=4995
  _globals['_USERSRESPONSE']._serialized_start=4997
  _globals['_USERSRESPONSE']._serialized_end=5056
  _globals['_USERISEXISTSREQUEST']._serialized_start=5058
  _globals['_USERISEXISTSREQUEST']._serialized_end=5094
  _globals['_USERISEXISTSRESPONSE']._serialized_start=5096
  _globals['_USERISEXISTSRESPONSE']._serialized_end=5136
  _globals['_COMPAREPASSWORDREQUEST']._serialized_start=5138
  _globals['_COMPAREPASSWORDREQUEST']._serialized_end=5195
  _globals['_COMPAREPASSWORDRESPONSE']._serialized_start=5197
  _globals['_COMPAREPASSWORDRESPONSE']._serialized_end=5241
  _globals['_COMPANY']._serialized_start=5244
  _globals['_COMPANY']._serialized_end=5456
  _globals['_CARDCOMPANY']._serialized_start=5459
  _globals['_CARDCOMPANY']._serialized_end=5602
  _globals['_CARD']._serialized_start=5605
  _globals['_CARD']._serialized_end=5738
  _globals['_WARD']._serialized_start=5741
  _globals['_WARD']._serialized_end=5903
  _globals['_DONATIONS']._serialized_start=5906
  _globals['_DONATIONS']._serialized_end=6043
  _globals['_CREATEUSERREQUEST']._serialized_start=6046
  _globals['_CREATEUSERREQUEST']._serialized_end=6301
  _globals['_CREATEUSERRESPONSE']._serialized_start=6304
  _globals['_CREATEUSERRESPONSE']._serialized_end=6591
  _globals['_ISROLEREQUEST']._serialized_start=6593
  _globals['_ISROLEREQUEST']._serialized_end=6634
  _globals['_ISROLERESPONSE']._serialized_start=6636
  _globals['_ISROLERESPONSE']._serialized_end=6671
  _globals['_DATABASESERVICE']._serialized_start=6674
  _globals['_DATABASESERVICE']._serialized_end=11068
# @@protoc_insertion_point(module_scope)
