from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Empty(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class UpdateUserCardRequest(_message.Message):
    __slots__ = ("id", "fullName", "number", "date", "cvv")
    ID_FIELD_NUMBER: _ClassVar[int]
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    NUMBER_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    CVV_FIELD_NUMBER: _ClassVar[int]
    id: int
    fullName: str
    number: str
    date: str
    cvv: int
    def __init__(self, id: _Optional[int] = ..., fullName: _Optional[str] = ..., number: _Optional[str] = ..., date: _Optional[str] = ..., cvv: _Optional[int] = ...) -> None: ...

class UpdateUserCardResponse(_message.Message):
    __slots__ = ("id", "fullName", "number", "date", "cvv", "userId", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    NUMBER_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    CVV_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    fullName: str
    number: str
    date: str
    cvv: int
    userId: int
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., fullName: _Optional[str] = ..., number: _Optional[str] = ..., date: _Optional[str] = ..., cvv: _Optional[int] = ..., userId: _Optional[int] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class AddCardToCompanyResponse(_message.Message):
    __slots__ = ("id", "title", "phone", "address", "site", "inn", "kpp", "okpo", "card", "userId", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    SITE_FIELD_NUMBER: _ClassVar[int]
    INN_FIELD_NUMBER: _ClassVar[int]
    KPP_FIELD_NUMBER: _ClassVar[int]
    OKPO_FIELD_NUMBER: _ClassVar[int]
    CARD_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    title: str
    phone: str
    address: str
    site: str
    inn: str
    kpp: str
    okpo: str
    card: CardCompany
    userId: int
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., title: _Optional[str] = ..., phone: _Optional[str] = ..., address: _Optional[str] = ..., site: _Optional[str] = ..., inn: _Optional[str] = ..., kpp: _Optional[str] = ..., okpo: _Optional[str] = ..., card: _Optional[_Union[CardCompany, _Mapping]] = ..., userId: _Optional[int] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class AddCardToCompanyRequest(_message.Message):
    __slots__ = ("card",)
    CARD_FIELD_NUMBER: _ClassVar[int]
    card: CreateCardCompanyRequest
    def __init__(self, card: _Optional[_Union[CreateCardCompanyRequest, _Mapping]] = ...) -> None: ...

class DeleteSessionByUserIdRequest(_message.Message):
    __slots__ = ("userId",)
    USERID_FIELD_NUMBER: _ClassVar[int]
    userId: int
    def __init__(self, userId: _Optional[int] = ...) -> None: ...

class DeleteSessionByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class DeleteSessionByModelRequest(_message.Message):
    __slots__ = ("id", "userId", "refreshToken", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    REFRESHTOKEN_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    userId: int
    refreshToken: str
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., userId: _Optional[int] = ..., refreshToken: _Optional[str] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class ChangeRefreshTokenByUserIdRequest(_message.Message):
    __slots__ = ("userId", "refreshToken")
    USERID_FIELD_NUMBER: _ClassVar[int]
    REFRESHTOKEN_FIELD_NUMBER: _ClassVar[int]
    userId: int
    refreshToken: str
    def __init__(self, userId: _Optional[int] = ..., refreshToken: _Optional[str] = ...) -> None: ...

class ChangeRefreshTokenByUserIdResponse(_message.Message):
    __slots__ = ("accessory",)
    ACCESSORY_FIELD_NUMBER: _ClassVar[int]
    accessory: bool
    def __init__(self, accessory: bool = ...) -> None: ...

class ChangeRefreshTokenByIdRequest(_message.Message):
    __slots__ = ("id", "refreshToken")
    ID_FIELD_NUMBER: _ClassVar[int]
    REFRESHTOKEN_FIELD_NUMBER: _ClassVar[int]
    id: int
    refreshToken: str
    def __init__(self, id: _Optional[int] = ..., refreshToken: _Optional[str] = ...) -> None: ...

class ChangeRefreshTokenByIdResponse(_message.Message):
    __slots__ = ("accessory",)
    ACCESSORY_FIELD_NUMBER: _ClassVar[int]
    accessory: bool
    def __init__(self, accessory: bool = ...) -> None: ...

class FindSessionsByUserIdRequest(_message.Message):
    __slots__ = ("userId",)
    USERID_FIELD_NUMBER: _ClassVar[int]
    userId: int
    def __init__(self, userId: _Optional[int] = ...) -> None: ...

class FindSessionsByUserIdResponse(_message.Message):
    __slots__ = ("id", "userId", "refreshToken", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    REFRESHTOKEN_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    userId: int
    refreshToken: str
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., userId: _Optional[int] = ..., refreshToken: _Optional[str] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class FindSessionsByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindSessionsByIdResponse(_message.Message):
    __slots__ = ("id", "userId", "refreshToken", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    REFRESHTOKEN_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    userId: int
    refreshToken: str
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., userId: _Optional[int] = ..., refreshToken: _Optional[str] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class SessionsResponse(_message.Message):
    __slots__ = ("sessions",)
    SESSIONS_FIELD_NUMBER: _ClassVar[int]
    sessions: _containers.RepeatedCompositeFieldContainer[CreateSessionResponse]
    def __init__(self, sessions: _Optional[_Iterable[_Union[CreateSessionResponse, _Mapping]]] = ...) -> None: ...

class CreateSessionRequest(_message.Message):
    __slots__ = ("userId", "refreshToken")
    USERID_FIELD_NUMBER: _ClassVar[int]
    REFRESHTOKEN_FIELD_NUMBER: _ClassVar[int]
    userId: int
    refreshToken: str
    def __init__(self, userId: _Optional[int] = ..., refreshToken: _Optional[str] = ...) -> None: ...

class CreateSessionResponse(_message.Message):
    __slots__ = ("id", "userId", "refreshToken", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    REFRESHTOKEN_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    userId: int
    refreshToken: str
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., userId: _Optional[int] = ..., refreshToken: _Optional[str] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class FindUserCompanyRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class HTTPCodes(_message.Message):
    __slots__ = ("code",)
    CODE_FIELD_NUMBER: _ClassVar[int]
    code: int
    def __init__(self, code: _Optional[int] = ...) -> None: ...

class AddCardToUserRequest(_message.Message):
    __slots__ = ("card",)
    CARD_FIELD_NUMBER: _ClassVar[int]
    card: CreateCardRequest
    def __init__(self, card: _Optional[_Union[CreateCardRequest, _Mapping]] = ...) -> None: ...

class UpdateDonationsRequest(_message.Message):
    __slots__ = ("id", "title", "amount", "wards", "userId", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    WARDS_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    title: str
    amount: float
    wards: Ward
    userId: int
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., title: _Optional[str] = ..., amount: _Optional[float] = ..., wards: _Optional[_Union[Ward, _Mapping]] = ..., userId: _Optional[int] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class DeleteDonationByModelRequest(_message.Message):
    __slots__ = ("id", "title", "amount", "wards", "userId", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    WARDS_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    title: str
    amount: float
    wards: Ward
    userId: int
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., title: _Optional[str] = ..., amount: _Optional[float] = ..., wards: _Optional[_Union[Ward, _Mapping]] = ..., userId: _Optional[int] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class AddCardToUserResponse(_message.Message):
    __slots__ = ("id", "email", "username", "password", "phone", "card", "role", "company", "type", "donations", "AvatarPath", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    CARD_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    COMPANY_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DONATIONS_FIELD_NUMBER: _ClassVar[int]
    AVATARPATH_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    email: str
    username: str
    password: str
    phone: str
    card: _containers.RepeatedCompositeFieldContainer[Card]
    role: str
    company: Company
    type: int
    donations: Donations
    AvatarPath: str
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., email: _Optional[str] = ..., username: _Optional[str] = ..., password: _Optional[str] = ..., phone: _Optional[str] = ..., card: _Optional[_Iterable[_Union[Card, _Mapping]]] = ..., role: _Optional[str] = ..., company: _Optional[_Union[Company, _Mapping]] = ..., type: _Optional[int] = ..., donations: _Optional[_Union[Donations, _Mapping]] = ..., AvatarPath: _Optional[str] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class CreateDonationsResponse(_message.Message):
    __slots__ = ("id", "title", "amount", "wards", "userId", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    WARDS_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    title: str
    amount: float
    wards: Ward
    userId: int
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., title: _Optional[str] = ..., amount: _Optional[float] = ..., wards: _Optional[_Union[Ward, _Mapping]] = ..., userId: _Optional[int] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class DeleteWardByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindWardByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class CreateWardRequest(_message.Message):
    __slots__ = ("title", "fullName", "want", "necessary", "donationId")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    WANT_FIELD_NUMBER: _ClassVar[int]
    NECESSARY_FIELD_NUMBER: _ClassVar[int]
    DONATIONID_FIELD_NUMBER: _ClassVar[int]
    title: str
    fullName: str
    want: str
    necessary: float
    donationId: int
    def __init__(self, title: _Optional[str] = ..., fullName: _Optional[str] = ..., want: _Optional[str] = ..., necessary: _Optional[float] = ..., donationId: _Optional[int] = ...) -> None: ...

class WardsResponse(_message.Message):
    __slots__ = ("wards",)
    WARDS_FIELD_NUMBER: _ClassVar[int]
    wards: _containers.RepeatedCompositeFieldContainer[Ward]
    def __init__(self, wards: _Optional[_Iterable[_Union[Ward, _Mapping]]] = ...) -> None: ...

class DeleteDonationByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindDonationByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindDonationWardsResponse(_message.Message):
    __slots__ = ("wards",)
    WARDS_FIELD_NUMBER: _ClassVar[int]
    wards: Ward
    def __init__(self, wards: _Optional[_Union[Ward, _Mapping]] = ...) -> None: ...

class FindDonationWardsRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class CreateDonationsRequest(_message.Message):
    __slots__ = ("title", "amount", "wards", "userId")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    WARDS_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    title: str
    amount: float
    wards: Ward
    userId: int
    def __init__(self, title: _Optional[str] = ..., amount: _Optional[float] = ..., wards: _Optional[_Union[Ward, _Mapping]] = ..., userId: _Optional[int] = ...) -> None: ...

class DonationsResponse(_message.Message):
    __slots__ = ("donations",)
    DONATIONS_FIELD_NUMBER: _ClassVar[int]
    donations: _containers.RepeatedCompositeFieldContainer[Donations]
    def __init__(self, donations: _Optional[_Iterable[_Union[Donations, _Mapping]]] = ...) -> None: ...

class DeleteCardCompanyByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindCardCompanyByIDRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class CreateCardCompanyRequest(_message.Message):
    __slots__ = ("fullName", "number", "date", "cvv", "companyId")
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    NUMBER_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    CVV_FIELD_NUMBER: _ClassVar[int]
    COMPANYID_FIELD_NUMBER: _ClassVar[int]
    fullName: str
    number: str
    date: str
    cvv: str
    companyId: int
    def __init__(self, fullName: _Optional[str] = ..., number: _Optional[str] = ..., date: _Optional[str] = ..., cvv: _Optional[str] = ..., companyId: _Optional[int] = ...) -> None: ...

class CardsCompaniesResponse(_message.Message):
    __slots__ = ("cards",)
    CARDS_FIELD_NUMBER: _ClassVar[int]
    cards: _containers.RepeatedCompositeFieldContainer[CardCompany]
    def __init__(self, cards: _Optional[_Iterable[_Union[CardCompany, _Mapping]]] = ...) -> None: ...

class DeleteCardByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindCardByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class CreateCardRequest(_message.Message):
    __slots__ = ("fullName", "number", "date", "cvv", "userId")
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    NUMBER_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    CVV_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    fullName: str
    number: str
    date: str
    cvv: int
    userId: int
    def __init__(self, fullName: _Optional[str] = ..., number: _Optional[str] = ..., date: _Optional[str] = ..., cvv: _Optional[int] = ..., userId: _Optional[int] = ...) -> None: ...

class CardsResponse(_message.Message):
    __slots__ = ("cards",)
    CARDS_FIELD_NUMBER: _ClassVar[int]
    cards: _containers.RepeatedCompositeFieldContainer[Card]
    def __init__(self, cards: _Optional[_Iterable[_Union[Card, _Mapping]]] = ...) -> None: ...

class UpdateCompanyRequest(_message.Message):
    __slots__ = ("company",)
    COMPANY_FIELD_NUMBER: _ClassVar[int]
    company: Company
    def __init__(self, company: _Optional[_Union[Company, _Mapping]] = ...) -> None: ...

class DeleteCompanyByModelRequest(_message.Message):
    __slots__ = ("company",)
    COMPANY_FIELD_NUMBER: _ClassVar[int]
    company: Company
    def __init__(self, company: _Optional[_Union[Company, _Mapping]] = ...) -> None: ...

class DeleteCompanyByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindCompanyCardRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindCompanyByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindCompanyByIdPhoneRequest(_message.Message):
    __slots__ = ("phone",)
    PHONE_FIELD_NUMBER: _ClassVar[int]
    phone: str
    def __init__(self, phone: _Optional[str] = ...) -> None: ...

class CreateCompanyRequest(_message.Message):
    __slots__ = ("title", "phone", "address", "site", "inn", "kpp", "okpo", "card", "userId")
    TITLE_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    SITE_FIELD_NUMBER: _ClassVar[int]
    INN_FIELD_NUMBER: _ClassVar[int]
    KPP_FIELD_NUMBER: _ClassVar[int]
    OKPO_FIELD_NUMBER: _ClassVar[int]
    CARD_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    title: str
    phone: str
    address: str
    site: str
    inn: str
    kpp: str
    okpo: str
    card: CardCompany
    userId: int
    def __init__(self, title: _Optional[str] = ..., phone: _Optional[str] = ..., address: _Optional[str] = ..., site: _Optional[str] = ..., inn: _Optional[str] = ..., kpp: _Optional[str] = ..., okpo: _Optional[str] = ..., card: _Optional[_Union[CardCompany, _Mapping]] = ..., userId: _Optional[int] = ...) -> None: ...

class CompaniesResponse(_message.Message):
    __slots__ = ("companies",)
    COMPANIES_FIELD_NUMBER: _ClassVar[int]
    companies: _containers.RepeatedCompositeFieldContainer[Company]
    def __init__(self, companies: _Optional[_Iterable[_Union[Company, _Mapping]]] = ...) -> None: ...

class UpdateUserRequest(_message.Message):
    __slots__ = ("id", "email", "username", "password", "phone", "card", "role", "company", "type", "donations", "AvatarPath", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    CARD_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    COMPANY_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DONATIONS_FIELD_NUMBER: _ClassVar[int]
    AVATARPATH_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    email: str
    username: str
    password: str
    phone: str
    card: _containers.RepeatedCompositeFieldContainer[Card]
    role: str
    company: Company
    type: int
    donations: _containers.RepeatedCompositeFieldContainer[Donations]
    AvatarPath: str
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., email: _Optional[str] = ..., username: _Optional[str] = ..., password: _Optional[str] = ..., phone: _Optional[str] = ..., card: _Optional[_Iterable[_Union[Card, _Mapping]]] = ..., role: _Optional[str] = ..., company: _Optional[_Union[Company, _Mapping]] = ..., type: _Optional[int] = ..., donations: _Optional[_Iterable[_Union[Donations, _Mapping]]] = ..., AvatarPath: _Optional[str] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class DeleteUserByModelRequest(_message.Message):
    __slots__ = ("user",)
    USER_FIELD_NUMBER: _ClassVar[int]
    user: CreateUserResponse
    def __init__(self, user: _Optional[_Union[CreateUserResponse, _Mapping]] = ...) -> None: ...

class DeleteUserByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindUserCardRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindUserCardResponse(_message.Message):
    __slots__ = ("cards",)
    CARDS_FIELD_NUMBER: _ClassVar[int]
    cards: _containers.RepeatedCompositeFieldContainer[Card]
    def __init__(self, cards: _Optional[_Iterable[_Union[Card, _Mapping]]] = ...) -> None: ...

class ChangeUserTypeRequest(_message.Message):
    __slots__ = ("id", "type")
    ID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    id: int
    type: int
    def __init__(self, id: _Optional[int] = ..., type: _Optional[int] = ...) -> None: ...

class ChangeUserTypeResponse(_message.Message):
    __slots__ = ("accessory",)
    ACCESSORY_FIELD_NUMBER: _ClassVar[int]
    accessory: bool
    def __init__(self, accessory: bool = ...) -> None: ...

class FindUserByPhoneRequest(_message.Message):
    __slots__ = ("phone",)
    PHONE_FIELD_NUMBER: _ClassVar[int]
    phone: str
    def __init__(self, phone: _Optional[str] = ...) -> None: ...

class FindUserByIdRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindUserDonationsRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...

class FindUserDonationsResponse(_message.Message):
    __slots__ = ("donations",)
    DONATIONS_FIELD_NUMBER: _ClassVar[int]
    donations: _containers.RepeatedCompositeFieldContainer[Donations]
    def __init__(self, donations: _Optional[_Iterable[_Union[Donations, _Mapping]]] = ...) -> None: ...

class FindUserByEmailRequest(_message.Message):
    __slots__ = ("email",)
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    email: str
    def __init__(self, email: _Optional[str] = ...) -> None: ...

class UsersResponse(_message.Message):
    __slots__ = ("users",)
    USERS_FIELD_NUMBER: _ClassVar[int]
    users: _containers.RepeatedCompositeFieldContainer[CreateUserResponse]
    def __init__(self, users: _Optional[_Iterable[_Union[CreateUserResponse, _Mapping]]] = ...) -> None: ...

class UserIsExistsRequest(_message.Message):
    __slots__ = ("phone",)
    PHONE_FIELD_NUMBER: _ClassVar[int]
    phone: str
    def __init__(self, phone: _Optional[str] = ...) -> None: ...

class UserIsExistsResponse(_message.Message):
    __slots__ = ("isExists",)
    ISEXISTS_FIELD_NUMBER: _ClassVar[int]
    isExists: bool
    def __init__(self, isExists: bool = ...) -> None: ...

class ComparePasswordRequest(_message.Message):
    __slots__ = ("phone", "password")
    PHONE_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    phone: str
    password: str
    def __init__(self, phone: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class ComparePasswordResponse(_message.Message):
    __slots__ = ("accessory",)
    ACCESSORY_FIELD_NUMBER: _ClassVar[int]
    accessory: bool
    def __init__(self, accessory: bool = ...) -> None: ...

class Company(_message.Message):
    __slots__ = ("id", "title", "phone", "address", "site", "inn", "kpp", "okpo", "card", "userId", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    SITE_FIELD_NUMBER: _ClassVar[int]
    INN_FIELD_NUMBER: _ClassVar[int]
    KPP_FIELD_NUMBER: _ClassVar[int]
    OKPO_FIELD_NUMBER: _ClassVar[int]
    CARD_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    title: str
    phone: str
    address: str
    site: str
    inn: str
    kpp: str
    okpo: str
    card: CardCompany
    userId: int
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., title: _Optional[str] = ..., phone: _Optional[str] = ..., address: _Optional[str] = ..., site: _Optional[str] = ..., inn: _Optional[str] = ..., kpp: _Optional[str] = ..., okpo: _Optional[str] = ..., card: _Optional[_Union[CardCompany, _Mapping]] = ..., userId: _Optional[int] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class CardCompany(_message.Message):
    __slots__ = ("id", "fullName", "number", "date", "cvv", "companyId", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    NUMBER_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    CVV_FIELD_NUMBER: _ClassVar[int]
    COMPANYID_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    fullName: str
    number: str
    date: str
    cvv: int
    companyId: int
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., fullName: _Optional[str] = ..., number: _Optional[str] = ..., date: _Optional[str] = ..., cvv: _Optional[int] = ..., companyId: _Optional[int] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class Card(_message.Message):
    __slots__ = ("id", "fullName", "number", "date", "cvv", "userId", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    NUMBER_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    CVV_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    fullName: str
    number: str
    date: str
    cvv: int
    userId: int
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., fullName: _Optional[str] = ..., number: _Optional[str] = ..., date: _Optional[str] = ..., cvv: _Optional[int] = ..., userId: _Optional[int] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class Ward(_message.Message):
    __slots__ = ("id", "title", "fullName", "want", "necessary", "donationId", "AvatarPath", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    WANT_FIELD_NUMBER: _ClassVar[int]
    NECESSARY_FIELD_NUMBER: _ClassVar[int]
    DONATIONID_FIELD_NUMBER: _ClassVar[int]
    AVATARPATH_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    title: str
    fullName: str
    want: str
    necessary: float
    donationId: int
    AvatarPath: str
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., title: _Optional[str] = ..., fullName: _Optional[str] = ..., want: _Optional[str] = ..., necessary: _Optional[float] = ..., donationId: _Optional[int] = ..., AvatarPath: _Optional[str] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class Donations(_message.Message):
    __slots__ = ("id", "title", "amount", "ward", "userId", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    WARD_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    title: str
    amount: float
    ward: Ward
    userId: int
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., title: _Optional[str] = ..., amount: _Optional[float] = ..., ward: _Optional[_Union[Ward, _Mapping]] = ..., userId: _Optional[int] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class CreateUserRequest(_message.Message):
    __slots__ = ("email", "username", "password", "phone", "card", "role", "company", "type", "donations")
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    CARD_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    COMPANY_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DONATIONS_FIELD_NUMBER: _ClassVar[int]
    email: str
    username: str
    password: str
    phone: str
    card: _containers.RepeatedCompositeFieldContainer[CreateCardRequest]
    role: str
    company: CreateCompanyRequest
    type: int
    donations: _containers.RepeatedCompositeFieldContainer[CreateDonationsRequest]
    def __init__(self, email: _Optional[str] = ..., username: _Optional[str] = ..., password: _Optional[str] = ..., phone: _Optional[str] = ..., card: _Optional[_Iterable[_Union[CreateCardRequest, _Mapping]]] = ..., role: _Optional[str] = ..., company: _Optional[_Union[CreateCompanyRequest, _Mapping]] = ..., type: _Optional[int] = ..., donations: _Optional[_Iterable[_Union[CreateDonationsRequest, _Mapping]]] = ...) -> None: ...

class CreateUserResponse(_message.Message):
    __slots__ = ("id", "email", "username", "password", "phone", "card", "role", "company", "type", "donations", "AvatarPath", "createdAt", "updatedAt")
    ID_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    PHONE_FIELD_NUMBER: _ClassVar[int]
    CARD_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    COMPANY_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DONATIONS_FIELD_NUMBER: _ClassVar[int]
    AVATARPATH_FIELD_NUMBER: _ClassVar[int]
    CREATEDAT_FIELD_NUMBER: _ClassVar[int]
    UPDATEDAT_FIELD_NUMBER: _ClassVar[int]
    id: int
    email: str
    username: str
    password: str
    phone: str
    card: _containers.RepeatedCompositeFieldContainer[Card]
    role: str
    company: Company
    type: int
    donations: _containers.RepeatedCompositeFieldContainer[Donations]
    AvatarPath: str
    createdAt: str
    updatedAt: str
    def __init__(self, id: _Optional[int] = ..., email: _Optional[str] = ..., username: _Optional[str] = ..., password: _Optional[str] = ..., phone: _Optional[str] = ..., card: _Optional[_Iterable[_Union[Card, _Mapping]]] = ..., role: _Optional[str] = ..., company: _Optional[_Union[Company, _Mapping]] = ..., type: _Optional[int] = ..., donations: _Optional[_Iterable[_Union[Donations, _Mapping]]] = ..., AvatarPath: _Optional[str] = ..., createdAt: _Optional[str] = ..., updatedAt: _Optional[str] = ...) -> None: ...

class IsRoleRequest(_message.Message):
    __slots__ = ("id", "role")
    ID_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    id: int
    role: str
    def __init__(self, id: _Optional[int] = ..., role: _Optional[str] = ...) -> None: ...

class IsRoleResponse(_message.Message):
    __slots__ = ("accessory",)
    ACCESSORY_FIELD_NUMBER: _ClassVar[int]
    accessory: bool
    def __init__(self, accessory: bool = ...) -> None: ...
