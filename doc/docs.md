# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [DatabaseService.proto](#DatabaseService-proto)
    - [AddCardToUserRequest](#service-AddCardToUserRequest)
    - [AddCardToUserResponse](#service-AddCardToUserResponse)
    - [Card](#service-Card)
    - [CardCompany](#service-CardCompany)
    - [CardsCompaniesResponse](#service-CardsCompaniesResponse)
    - [CardsResponse](#service-CardsResponse)
    - [ChangeRefreshTokenByIdRequest](#service-ChangeRefreshTokenByIdRequest)
    - [ChangeRefreshTokenByIdResponse](#service-ChangeRefreshTokenByIdResponse)
    - [ChangeRefreshTokenByUserIdRequest](#service-ChangeRefreshTokenByUserIdRequest)
    - [ChangeRefreshTokenByUserIdResponse](#service-ChangeRefreshTokenByUserIdResponse)
    - [ChangeUserTypeRequest](#service-ChangeUserTypeRequest)
    - [ChangeUserTypeResponse](#service-ChangeUserTypeResponse)
    - [CompaniesResponse](#service-CompaniesResponse)
    - [Company](#service-Company)
    - [ComparePasswordRequest](#service-ComparePasswordRequest)
    - [ComparePasswordResponse](#service-ComparePasswordResponse)
    - [CreateCardCompanyRequest](#service-CreateCardCompanyRequest)
    - [CreateCardRequest](#service-CreateCardRequest)
    - [CreateCompanyRequest](#service-CreateCompanyRequest)
    - [CreateDonationsRequest](#service-CreateDonationsRequest)
    - [CreateDonationsResponse](#service-CreateDonationsResponse)
    - [CreateSessionRequest](#service-CreateSessionRequest)
    - [CreateSessionResponse](#service-CreateSessionResponse)
    - [CreateUserRequest](#service-CreateUserRequest)
    - [CreateUserResponse](#service-CreateUserResponse)
    - [CreateWardRequest](#service-CreateWardRequest)
    - [DeleteCardByIdRequest](#service-DeleteCardByIdRequest)
    - [DeleteCardCompanyByIdRequest](#service-DeleteCardCompanyByIdRequest)
    - [DeleteCompanyByIdRequest](#service-DeleteCompanyByIdRequest)
    - [DeleteCompanyByModelRequest](#service-DeleteCompanyByModelRequest)
    - [DeleteDonationByIdRequest](#service-DeleteDonationByIdRequest)
    - [DeleteDonationByModelRequest](#service-DeleteDonationByModelRequest)
    - [DeleteSessionByIdRequest](#service-DeleteSessionByIdRequest)
    - [DeleteSessionByModelRequest](#service-DeleteSessionByModelRequest)
    - [DeleteSessionByUserIdRequest](#service-DeleteSessionByUserIdRequest)
    - [DeleteUserByIdRequest](#service-DeleteUserByIdRequest)
    - [DeleteUserByModelRequest](#service-DeleteUserByModelRequest)
    - [DeleteWardByIdRequest](#service-DeleteWardByIdRequest)
    - [Donations](#service-Donations)
    - [DonationsResponse](#service-DonationsResponse)
    - [Empty](#service-Empty)
    - [FindCardByIdRequest](#service-FindCardByIdRequest)
    - [FindCardCompanyByIDRequest](#service-FindCardCompanyByIDRequest)
    - [FindCompanyByIdPhoneRequest](#service-FindCompanyByIdPhoneRequest)
    - [FindCompanyByIdRequest](#service-FindCompanyByIdRequest)
    - [FindCompanyCardRequest](#service-FindCompanyCardRequest)
    - [FindDonationByIdRequest](#service-FindDonationByIdRequest)
    - [FindDonationWardsRequest](#service-FindDonationWardsRequest)
    - [FindDonationWardsResponse](#service-FindDonationWardsResponse)
    - [FindSessionsByIdRequest](#service-FindSessionsByIdRequest)
    - [FindSessionsByIdResponse](#service-FindSessionsByIdResponse)
    - [FindSessionsByUserIdRequest](#service-FindSessionsByUserIdRequest)
    - [FindSessionsByUserIdResponse](#service-FindSessionsByUserIdResponse)
    - [FindUserByEmailRequest](#service-FindUserByEmailRequest)
    - [FindUserByIdRequest](#service-FindUserByIdRequest)
    - [FindUserByPhoneRequest](#service-FindUserByPhoneRequest)
    - [FindUserCardRequest](#service-FindUserCardRequest)
    - [FindUserCardResponse](#service-FindUserCardResponse)
    - [FindUserCompanyRequest](#service-FindUserCompanyRequest)
    - [FindUserDonationsRequest](#service-FindUserDonationsRequest)
    - [FindUserDonationsResponse](#service-FindUserDonationsResponse)
    - [FindWardByIdRequest](#service-FindWardByIdRequest)
    - [HTTPCodes](#service-HTTPCodes)
    - [IsRoleRequest](#service-IsRoleRequest)
    - [IsRoleResponse](#service-IsRoleResponse)
    - [SessionsResponse](#service-SessionsResponse)
    - [UpdateCompanyRequest](#service-UpdateCompanyRequest)
    - [UpdateDonationsRequest](#service-UpdateDonationsRequest)
    - [UpdateUserRequest](#service-UpdateUserRequest)
    - [UserIsExistsRequest](#service-UserIsExistsRequest)
    - [UserIsExistsResponse](#service-UserIsExistsResponse)
    - [UsersResponse](#service-UsersResponse)
    - [Ward](#service-Ward)
    - [WardsResponse](#service-WardsResponse)
  
    - [DatabaseService](#service-DatabaseService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="DatabaseService-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## DatabaseService.proto



<a name="service-AddCardToUserRequest"></a>

### AddCardToUserRequest

Запрос на добавление банковской карты пользователю


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card | [CreateCardRequest](#service-CreateCardRequest) |  | Банковская карта |






<a name="service-AddCardToUserResponse"></a>

### AddCardToUserResponse

Ответ на запрос добавления банковской карты пользователю


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя в базе данных |
| email | [string](#string) |  | Email пользователя |
| username | [string](#string) |  | Имя (никнейм) пользователя |
| password | [string](#string) |  | Пароль пользователя (закодирован в MD5Hash) |
| phone | [string](#string) |  | Номер телефона пользователя |
| card | [Card](#service-Card) | repeated | Массив банковских карт пользователя |
| role | [string](#string) |  | Роль пользователя |
| company | [Company](#service-Company) |  | Компания пользователя, если он является юридическим лицом (type = 1) |
| type | [uint64](#uint64) |  | Тип пользователя (0 - физическое лицо, 1 - юридическое лицо) |
| donations | [Donations](#service-Donations) |  | Массив пожертвований пользователя |
| AvatarPath | [string](#string) |  | Локальный путь к аватару пользователя |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-Card"></a>

### Card

Банковская карта для пользователей


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID банковской карты в базе данных |
| fullName | [string](#string) |  | ФИО с банковской карты |
| number | [string](#string) |  | Номер карты |
| date | [string](#string) |  | Дата до которой активна карта |
| cvv | [uint64](#uint64) |  | CVV код карты |
| userId | [uint64](#uint64) |  | ID пользователя которому принадлежит данная карта |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-CardCompany"></a>

### CardCompany

Банковская карта компании


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID банковской карты в базе данных |
| fullName | [string](#string) |  | ФИО с банковской карты |
| number | [string](#string) |  | Номер карты |
| date | [string](#string) |  | Дата до которой активна карта |
| cvv | [uint64](#uint64) |  | CVV код карты |
| companyId | [uint64](#uint64) |  | ID компании которой принадлежит данная карта |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-CardsCompaniesResponse"></a>

### CardsCompaniesResponse

Запрос на получение всех банковских карт компаний


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cards | [CardCompany](#service-CardCompany) | repeated | Массив банковских карт компаний |






<a name="service-CardsResponse"></a>

### CardsResponse

Ответ на запрос получения всех банковских карт пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cards | [Card](#service-Card) | repeated | Массив банковских карт пользователя |






<a name="service-ChangeRefreshTokenByIdRequest"></a>

### ChangeRefreshTokenByIdRequest

Запрос на обновление Refresh Token по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID сессии |
| refreshToken | [string](#string) |  | Refresh Token |






<a name="service-ChangeRefreshTokenByIdResponse"></a>

### ChangeRefreshTokenByIdResponse

Ответ на запрос обновления Refresh Token по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accessory | [bool](#bool) |  | Успешность операции обновления Refresh Token |






<a name="service-ChangeRefreshTokenByUserIdRequest"></a>

### ChangeRefreshTokenByUserIdRequest

Запрос на обновление Refresh Token по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userId | [uint64](#uint64) |  | UserID(ID пользователя в БД) пользователя которому принадлежит сессия |
| refreshToken | [string](#string) |  | Refresh Token |






<a name="service-ChangeRefreshTokenByUserIdResponse"></a>

### ChangeRefreshTokenByUserIdResponse

Ответ на запрос обновления Refresh Token по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accessory | [bool](#bool) |  | Успешность операции обновления Refresh Token |






<a name="service-ChangeUserTypeRequest"></a>

### ChangeUserTypeRequest

Запрос на изменение типа пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя |
| type | [uint64](#uint64) |  | Тип пользователя на который меняем (0 - физическое лицо, 1 - юридическое лицо) |






<a name="service-ChangeUserTypeResponse"></a>

### ChangeUserTypeResponse

Ответ на запрос изменений типа пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accessory | [bool](#bool) |  | Успешность операции изменения типа (true/false) |






<a name="service-CompaniesResponse"></a>

### CompaniesResponse

Ответ на запрос извлечения всех компаний


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| companies | [Company](#service-Company) | repeated | Массив компаний |






<a name="service-Company"></a>

### Company

Компания


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID компании в базе данных |
| title | [string](#string) |  | Название компании |
| phone | [string](#string) |  | Номер телефона компании |
| address | [string](#string) |  | Адрес офиса компании |
| site | [string](#string) |  | Сайт компании |
| inn | [string](#string) |  | ИНН юридического лица |
| kpp | [string](#string) |  | КПП юридического лица |
| okpo | [string](#string) |  | ОКПО предприятия/организации |
| card | [CardCompany](#service-CardCompany) |  | Банковская карта компании |
| userId | [uint64](#uint64) |  | ID пользователя к которому относится данная компания |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-ComparePasswordRequest"></a>

### ComparePasswordRequest

Запрос на сравнение пароля


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | Номер телефона пользователя, чей пароль будем искать для сравнения |
| password | [string](#string) |  | Пароль который будем сравнивать с тем, что есть в базе данных |






<a name="service-ComparePasswordResponse"></a>

### ComparePasswordResponse

Ответ на запрос сравнения пароля


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accessory | [bool](#bool) |  | Совпадает ли пароль (true/false) |






<a name="service-CreateCardCompanyRequest"></a>

### CreateCardCompanyRequest

Запрос на создание банковской карты для компании


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fullName | [string](#string) |  | ФИО с банковской карты |
| number | [string](#string) |  | Номер банковской карты |
| date | [string](#string) |  | Дата до которой активна карта |
| cvv | [string](#string) |  | CVV код банковской карты |
| companyId | [uint64](#uint64) |  | ID компании которой принадлежит данная карта |






<a name="service-CreateCardRequest"></a>

### CreateCardRequest

Запрос на создание банковской карты для пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fullName | [string](#string) |  | ФИО с банковской карты |
| number | [string](#string) |  | Номер карты |
| date | [string](#string) |  | Дата до которой активна карта |
| cvv | [uint64](#uint64) |  | CVV код карты |
| userId | [uint64](#uint64) |  | ID пользователя кому принадлежит карта |






<a name="service-CreateCompanyRequest"></a>

### CreateCompanyRequest

Запрос на создание компании


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | Название компании |
| phone | [string](#string) |  | Номер телефона компании |
| address | [string](#string) |  | Адрес офиса компании |
| site | [string](#string) |  | Сайт компании |
| inn | [string](#string) |  | ИНН юридического лица |
| kpp | [string](#string) |  | КПП юридического лица |
| okpo | [string](#string) |  | ОКПО предприятия/организации |
| card | [CardCompany](#service-CardCompany) |  | Банковская карта компании |
| userId | [uint64](#uint64) |  | ID пользователя к которому относится данная компания |






<a name="service-CreateDonationsRequest"></a>

### CreateDonationsRequest

Запрос на создание пожертвования


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | Произвольное название пожертвования, примерно так: &#34;На лекарства&#34; |
| amount | [float](#float) |  | Сумма пожертвования |
| wards | [Ward](#service-Ward) |  | Подопечный данного пожертвования |
| userId | [uint64](#uint64) |  | ID пользователя, чье это пожертвование |






<a name="service-CreateDonationsResponse"></a>

### CreateDonationsResponse

Ответ на запрос создания пожертвования


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID записи в базе данных |
| title | [string](#string) |  | Название пожертвования (например &#34;На лекарства&#34;) |
| amount | [float](#float) |  | Сумма пожертвования |
| wards | [Ward](#service-Ward) |  | Подопечный этого пожертвования |
| userId | [uint64](#uint64) |  | ID пользователя, которому принадлежит пожертвование |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-CreateSessionRequest"></a>

### CreateSessionRequest

Запрос на создание сессии


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userId | [uint64](#uint64) |  | ID пользователя, кому принадлежит сессия |
| refreshToken | [string](#string) |  | Refresh token |






<a name="service-CreateSessionResponse"></a>

### CreateSessionResponse

Ответ на запрос создания сессии


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID сессии в базе данных |
| userId | [uint64](#uint64) |  | ID пользователя, кому принадлежит сессия |
| refreshToken | [string](#string) |  | Refresh token |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-CreateUserRequest"></a>

### CreateUserRequest

Запрос на создание пользователя (Поля &lt;b&gt;card&lt;/b&gt;, &lt;b&gt;company&lt;/b&gt;, &lt;b&gt;donations&lt;/b&gt; необязательные)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email пользователя (обязательно проверяем на формат email@email.com) |
| username | [string](#string) |  | Имя (никнейм) пользователя |
| password | [string](#string) |  | Пароль пользователя (необходимо дополнительно валидировать корректность пароля) |
| phone | [string](#string) |  | Номер телефона (необходимо дополнительно валидировать формат) |
| card | [CreateCardRequest](#service-CreateCardRequest) | repeated | Банковские карты пользователя (если влаживаем в этот запрос, то userId не указываем) |
| role | [string](#string) |  | Роль пользователя |
| company | [CreateCompanyRequest](#service-CreateCompanyRequest) |  | Компания пользователя, если он является юридическим лицом |
| type | [uint64](#uint64) |  | Тип пользователя, 0 - физическое лицо, 1 - юридическое лицо |
| donations | [CreateDonationsRequest](#service-CreateDonationsRequest) | repeated | Пожертвования пользователя |






<a name="service-CreateUserResponse"></a>

### CreateUserResponse

Ответ на запрос создания пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя |
| email | [string](#string) |  | Email пользователя |
| username | [string](#string) |  | Имя (никнейм) пользователя |
| password | [string](#string) |  | Пароль пользователя (закодирован в MD5Hash) |
| phone | [string](#string) |  | Номер телефона пользователя |
| card | [Card](#service-Card) | repeated | Банковские карты пользователя |
| role | [string](#string) |  | Роль пользователя |
| company | [Company](#service-Company) |  | Компания пользователя, если он является юридическим лицом |
| type | [uint64](#uint64) |  | Тип пользователя, 0 - физическое лицо, 1 - юридическое лицо |
| donations | [Donations](#service-Donations) | repeated | Пожертвования пользователя |
| AvatarPath | [string](#string) |  | Локальный путь к аватару пользователя |
| createdAt | [string](#string) |  | Дата создания сущности записи в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-CreateWardRequest"></a>

### CreateWardRequest

Запрос на создание подопечного


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | Дополнительный текст к подопечному |
| fullName | [string](#string) |  | Полное имя подопечного |
| want | [string](#string) |  | Необходимость подопечного (то в чем он нуждается, например &#34;Лекарства&#34;) |
| necessary | [float](#float) |  | Необходимая сумма денег на необходимость |
| donationId | [uint64](#uint64) |  | ID пожертвования |






<a name="service-DeleteCardByIdRequest"></a>

### DeleteCardByIdRequest

Запрос на удаление банковской карты пользователя по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID банковской карты |






<a name="service-DeleteCardCompanyByIdRequest"></a>

### DeleteCardCompanyByIdRequest

Удаление банковской карты компании по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID банковской карты компании, которую нужно удалить |






<a name="service-DeleteCompanyByIdRequest"></a>

### DeleteCompanyByIdRequest

Запрос на удаление компании по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID компании которую будем удалять |






<a name="service-DeleteCompanyByModelRequest"></a>

### DeleteCompanyByModelRequest

Удаление компании по модели


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| company | [Company](#service-Company) |  | Полная сущность модели для удаления |






<a name="service-DeleteDonationByIdRequest"></a>

### DeleteDonationByIdRequest

Запрос на удаление пожертвования по его ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пожертвования которое будем удалять |






<a name="service-DeleteDonationByModelRequest"></a>

### DeleteDonationByModelRequest

Запрос на удаление пожертвования по модели


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пожертвования в базе данных |
| title | [string](#string) |  | Название пожертвования (например &#34;На лекарства&#34;) |
| amount | [float](#float) |  | Сумма пожертвования |
| wards | [Ward](#service-Ward) |  | Подопечный этого пожертвования |
| userId | [uint64](#uint64) |  | ID пользователя, которому принадлежит пожертвование |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-DeleteSessionByIdRequest"></a>

### DeleteSessionByIdRequest

Запрос на удаление сессии по ее ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID сессии которую будем удалять |






<a name="service-DeleteSessionByModelRequest"></a>

### DeleteSessionByModelRequest

Запрос на удаление сессии по сущности (например полученную с помощью FindSessionById)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID сессии в базе данных |
| userId | [uint64](#uint64) |  | ID пользователя, кому принадлежит сессия |
| refreshToken | [string](#string) |  | Refresh token |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-DeleteSessionByUserIdRequest"></a>

### DeleteSessionByUserIdRequest

Запрос на удаление сессии по ее ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userId | [uint64](#uint64) |  | UserID(ID пользователя в БД) пользователя, сессию которую будем удалять |






<a name="service-DeleteUserByIdRequest"></a>

### DeleteUserByIdRequest

Запрос на удаление пользователя по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя для удаления |






<a name="service-DeleteUserByModelRequest"></a>

### DeleteUserByModelRequest

Запрос на удаление пользователя по сущности


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [CreateUserResponse](#service-CreateUserResponse) |  | Сущность пользователя для удаления (полная модель пользователя) |






<a name="service-DeleteWardByIdRequest"></a>

### DeleteWardByIdRequest

Удаление подопечного по его ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID подопечного которого будем удалять |






<a name="service-Donations"></a>

### Donations

Пожертвования


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пожертвования в базе данных |
| title | [string](#string) |  | Название пожертвования (например &#34;На лекарства&#34;) |
| amount | [float](#float) |  | Сумма пожертвования |
| ward | [Ward](#service-Ward) |  | Подопечный этого пожертвования |
| userId | [uint64](#uint64) |  | ID пользователя, которому принадлежит пожертвование |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-DonationsResponse"></a>

### DonationsResponse

Запрос на получение всех пожертвований


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| donations | [Donations](#service-Donations) | repeated | Массив пожертвований |






<a name="service-Empty"></a>

### Empty

Пустая структура, необходимая для запросов без входящих аргументов






<a name="service-FindCardByIdRequest"></a>

### FindCardByIdRequest

Запрос на поиск банковской карты пользователя по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID банковской карты |






<a name="service-FindCardCompanyByIDRequest"></a>

### FindCardCompanyByIDRequest

Запрос на извлечение банковских карт компаний определенной компании


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID компании, банковские карты которой нужно извлечь |






<a name="service-FindCompanyByIdPhoneRequest"></a>

### FindCompanyByIdPhoneRequest

Запрос на поиск компании по номеру телефона


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | Номер телефона компании по которому будем искать |






<a name="service-FindCompanyByIdRequest"></a>

### FindCompanyByIdRequest

Запрос на поиск компании по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID по которому ищем компанию |






<a name="service-FindCompanyCardRequest"></a>

### FindCompanyCardRequest

Запрос на поиск компании по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID компании которую будем искать |






<a name="service-FindDonationByIdRequest"></a>

### FindDonationByIdRequest

Запрос на поиск пожертвования по его ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пожертвования которое будем искать |






<a name="service-FindDonationWardsRequest"></a>

### FindDonationWardsRequest

Запрос на извлечение подопечного пожертвования по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пожертвования |






<a name="service-FindDonationWardsResponse"></a>

### FindDonationWardsResponse

Ответ на запрос извлечения подопечного пожертвования


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wards | [Ward](#service-Ward) |  | Подопечный пожертвования |






<a name="service-FindSessionsByIdRequest"></a>

### FindSessionsByIdRequest

Запрос на поиск сессии по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID сессии которую необходимо найти |






<a name="service-FindSessionsByIdResponse"></a>

### FindSessionsByIdResponse

Ответ на запрос поиска сессии по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID сессии в базе данных |
| userId | [uint64](#uint64) |  | ID пользователя, кому принадлежит сессия |
| refreshToken | [string](#string) |  | Refresh token |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-FindSessionsByUserIdRequest"></a>

### FindSessionsByUserIdRequest

Запрос на поиск сессии по UserID(ID пользователя в БД)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userId | [uint64](#uint64) |  | UserID(ID пользователя в БД) чью сессию необходимо найти |






<a name="service-FindSessionsByUserIdResponse"></a>

### FindSessionsByUserIdResponse

Ответ на запрос поиска сессии по UserID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID сессии в базе данных |
| userId | [uint64](#uint64) |  | ID пользователя, кому принадлежит сессия |
| refreshToken | [string](#string) |  | Refresh token |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-FindUserByEmailRequest"></a>

### FindUserByEmailRequest

Запрос на поиск пользователя по email


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email пользователя для поиска |






<a name="service-FindUserByIdRequest"></a>

### FindUserByIdRequest

Запрос на поиск пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя которого будем искать |






<a name="service-FindUserByPhoneRequest"></a>

### FindUserByPhoneRequest

Запрос на поиск пользователя по номеру телефона


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | Номер телефона по которому будет искать пользователя |






<a name="service-FindUserCardRequest"></a>

### FindUserCardRequest

Запрос на извлечение банковских карт пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя чьи банковские карты нужно извлечь |






<a name="service-FindUserCardResponse"></a>

### FindUserCardResponse

Ответ на запрос извлечения банковских карт пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cards | [Card](#service-Card) | repeated | Массив банковских карт |






<a name="service-FindUserCompanyRequest"></a>

### FindUserCompanyRequest

Запрос на извлечение компании пользователя (если его поле type = 1)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя чью компанию нужно найти |






<a name="service-FindUserDonationsRequest"></a>

### FindUserDonationsRequest

Запрос на извлечение пожертвований пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя чьи пожертвования нужно извлечь |






<a name="service-FindUserDonationsResponse"></a>

### FindUserDonationsResponse

Ответ на запрос извлечения пожертвований пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| donations | [Donations](#service-Donations) | repeated | Массив пожертвований пользователя |






<a name="service-FindWardByIdRequest"></a>

### FindWardByIdRequest

Запрос на поиск подопечного по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID подопечного которого будем искать |






<a name="service-HTTPCodes"></a>

### HTTPCodes

Стандартные HTTP ответы на запросы


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | HTTP код ответа |






<a name="service-IsRoleRequest"></a>

### IsRoleRequest

Запрос на проверку принадлежности пользователя к роли


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя |
| role | [string](#string) |  | Роль пользователя которую мы ожиданием |






<a name="service-IsRoleResponse"></a>

### IsRoleResponse

Ответ на запрос проверки принадлежности пользователя к роли


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accessory | [bool](#bool) |  | Принадлежность к роли (true/false) |






<a name="service-SessionsResponse"></a>

### SessionsResponse

Ответ на запрос списка всех сессий в базе данных


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sessions | [CreateSessionResponse](#service-CreateSessionResponse) | repeated | Массив сессий |






<a name="service-UpdateCompanyRequest"></a>

### UpdateCompanyRequest

Запрос на обновление компании


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| company | [Company](#service-Company) |  | Компания которую обновляем |






<a name="service-UpdateDonationsRequest"></a>

### UpdateDonationsRequest

Запрос на обновление пожертвования


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пожертвования в базе данных |
| title | [string](#string) |  | Название пожертвования (например &#34;На лекарства&#34;) |
| amount | [float](#float) |  | Сумма пожертвования |
| wards | [Ward](#service-Ward) |  | Подопечный этого пожертвования |
| userId | [uint64](#uint64) |  | ID пользователя, которому принадлежит пожертвование |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-UpdateUserRequest"></a>

### UpdateUserRequest

Запрос на обновление сущности пользователя


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID пользователя в базе данных |
| email | [string](#string) |  | Email пользователя |
| username | [string](#string) |  | Имя (никнейм) пользователя |
| password | [string](#string) |  | Пароль пользователя (закодирован в MD5Hash) |
| phone | [string](#string) |  | Номер телефона пользователя |
| card | [Card](#service-Card) | repeated | Массив банковских карт пользователя |
| role | [string](#string) |  | Роль пользователя |
| company | [Company](#service-Company) |  | Компания пользователя, если он является юридическим лицом (type = 1) |
| type | [uint64](#uint64) |  | Тип пользователя (0 - физическое лицо, 1 - юридическое лицо) |
| donations | [Donations](#service-Donations) | repeated | Массив пожертвований пользователя |
| AvatarPath | [string](#string) |  | Локальный путь к аватару пользователя |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-UserIsExistsRequest"></a>

### UserIsExistsRequest

Запрос на проверку существования пользователя в базе данных


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  | Номер телефона пользователя, по которому будем искать сущность |






<a name="service-UserIsExistsResponse"></a>

### UserIsExistsResponse

Ответ на запрос проверки существования пользователя в базе данных


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isExists | [bool](#bool) |  | Существует ли пользователь в базе данных (true/false) |






<a name="service-UsersResponse"></a>

### UsersResponse

Ответ на запрос списка всех пользователей в базе данных


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [CreateUserResponse](#service-CreateUserResponse) | repeated | Массив пользователей |






<a name="service-Ward"></a>

### Ward

Подопечные


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | ID подопечного в базе данных |
| title | [string](#string) |  | Дополнительный текст к подопечному |
| fullName | [string](#string) |  | Полное имя подопечного |
| want | [string](#string) |  | Необходимость подопечного (то в чем он нуждается, например &#34;Лекарства&#34;) |
| necessary | [float](#float) |  | Необходимая сумма денег на необходимость |
| donationId | [uint64](#uint64) |  | ID пожертвования к которому относится данный подопечный |
| AvatarPath | [string](#string) |  | Локальный путь к аватару подопечного |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-WardsResponse"></a>

### WardsResponse

Ответ на запрос получения всех подопечных


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wards | [Ward](#service-Ward) | repeated | Массив подопечных |





 

 

 


<a name="service-DatabaseService"></a>

### DatabaseService

Доступные API методы

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [CreateUserRequest](#service-CreateUserRequest) | [CreateUserResponse](#service-CreateUserResponse) |  Создание нового пользователя |
| Users | [Empty](#service-Empty) | [UsersResponse](#service-UsersResponse) |  Список всех пользователей |
| IsRole | [IsRoleRequest](#service-IsRoleRequest) | [IsRoleResponse](#service-IsRoleResponse) |  Проверяет принадлежность пользователя к необходимой роли |
| ComparePassword | [ComparePasswordRequest](#service-ComparePasswordRequest) | [ComparePasswordResponse](#service-ComparePasswordResponse) |  Сравнивает введенный пароль пользователя с тем, что находится в базе данных |
| UserIsExists | [UserIsExistsRequest](#service-UserIsExistsRequest) | [UserIsExistsResponse](#service-UserIsExistsResponse) |  Проверка на существования пользователя в базе данных (проверяет по номеру телефона) |
| FindUserById | [FindUserByIdRequest](#service-FindUserByIdRequest) | [CreateUserResponse](#service-CreateUserResponse) |  Поиск пользователя по его ID |
| FindUserByEmail | [FindUserByEmailRequest](#service-FindUserByEmailRequest) | [CreateUserResponse](#service-CreateUserResponse) |  Поиск пользователя по его email |
| FindUserByPhone | [FindUserByPhoneRequest](#service-FindUserByPhoneRequest) | [CreateUserResponse](#service-CreateUserResponse) |  Поиск пользователя по номеру телефона |
| ChangeUserType | [ChangeUserTypeRequest](#service-ChangeUserTypeRequest) | [ChangeUserTypeResponse](#service-ChangeUserTypeResponse) |  Изменение типа пользователя (0 - физическое лицо, 1 - юридическое лицо) |
| FindUserCompany | [FindUserCompanyRequest](#service-FindUserCompanyRequest) | [Company](#service-Company) |  Поиск компании определенного пользователя |
| FindUserDonations | [FindUserDonationsRequest](#service-FindUserDonationsRequest) | [FindUserDonationsResponse](#service-FindUserDonationsResponse) |  Поиск пожертвований определенного пользователя |
| FindUserCard | [FindUserCardRequest](#service-FindUserCardRequest) | [FindUserCardResponse](#service-FindUserCardResponse) |  Поиск банковских карт определенного пользователя |
| AddCardToUser | [AddCardToUserRequest](#service-AddCardToUserRequest) | [AddCardToUserResponse](#service-AddCardToUserResponse) |  Добавляет банковскую карту пользователю |
| DeleteUserByModel | [DeleteUserByModelRequest](#service-DeleteUserByModelRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаление пользователя по полной модели (например полученной с помощью FindUserById) |
| DeleteUserById | [DeleteUserByIdRequest](#service-DeleteUserByIdRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаляет пользователя по его ID |
| UpdateUser | [UpdateUserRequest](#service-UpdateUserRequest) | [CreateUserResponse](#service-CreateUserResponse) |  Обновление пользователя |
| Companies | [Empty](#service-Empty) | [CompaniesResponse](#service-CompaniesResponse) |  Список всех компаний |
| CreateCompany | [CreateCompanyRequest](#service-CreateCompanyRequest) | [Company](#service-Company) |  Создание компании |
| FindCompanyById | [FindCompanyByIdRequest](#service-FindCompanyByIdRequest) | [Company](#service-Company) |  Поиск компании по ее ID |
| FindCompanyByPhone | [FindCompanyByIdPhoneRequest](#service-FindCompanyByIdPhoneRequest) | [Company](#service-Company) |  Поиск компании по номеру телефона |
| FindCompanyCard | [FindCompanyCardRequest](#service-FindCompanyCardRequest) | [CardCompany](#service-CardCompany) |  Поиск банковской карты компании |
| DeleteCompanyByModel | [DeleteCompanyByModelRequest](#service-DeleteCompanyByModelRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаление компании по модели (например полученной с помощью FindCompanyById) |
| DeleteCompanyById | [DeleteCompanyByIdRequest](#service-DeleteCompanyByIdRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаление компании по ее ID |
| UpdateCompany | [UpdateCompanyRequest](#service-UpdateCompanyRequest) | [HTTPCodes](#service-HTTPCodes) |  Обновление компании |
| Cards | [Empty](#service-Empty) | [CardsResponse](#service-CardsResponse) |  Поиск всех банковских карт пользователей |
| CreateCard | [CreateCardRequest](#service-CreateCardRequest) | [Card](#service-Card) |  Создание банковской карты |
| FindCardById | [FindCardByIdRequest](#service-FindCardByIdRequest) | [Card](#service-Card) |  Поиск банковской карты по ее ID |
| DeleteCardByModel | [Card](#service-Card) | [HTTPCodes](#service-HTTPCodes) |  Удаление банковской карты по модели (например полученной с помощью FindCardById) |
| DeleteCardById | [DeleteCardByIdRequest](#service-DeleteCardByIdRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаление банковской карты по ID |
| UpdateCard | [Card](#service-Card) | [Card](#service-Card) |  Обновление банковской карты |
| CardsCompanies | [Empty](#service-Empty) | [CardsCompaniesResponse](#service-CardsCompaniesResponse) |  Поиск всех банковских карт компаний |
| CreateCardCompany | [CreateCardCompanyRequest](#service-CreateCardCompanyRequest) | [CardCompany](#service-CardCompany) |  Создание банковской карты компании |
| FindCardCompanyByID | [FindCardCompanyByIDRequest](#service-FindCardCompanyByIDRequest) | [CardCompany](#service-CardCompany) |  Поиск банковской карты компании по ID |
| DeleteCardCompanyByModel | [CardCompany](#service-CardCompany) | [HTTPCodes](#service-HTTPCodes) |  Удаление банковской карты по модели (например полученной с помощью FindCardCompanyByID) |
| DeleteCardCompanyById | [DeleteCardCompanyByIdRequest](#service-DeleteCardCompanyByIdRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаление банковской карты компании по ID |
| UpdateCardCompany | [CardCompany](#service-CardCompany) | [CardCompany](#service-CardCompany) |  Обновление банковской карты компании |
| Donations | [Empty](#service-Empty) | [DonationsResponse](#service-DonationsResponse) |  Список всех пожертвований |
| CreateDonations | [CreateDonationsRequest](#service-CreateDonationsRequest) | [CreateDonationsResponse](#service-CreateDonationsResponse) |  Создание пожертвования |
| FindDonationWards | [FindDonationWardsRequest](#service-FindDonationWardsRequest) | [FindDonationWardsResponse](#service-FindDonationWardsResponse) |  Поиск подопечных по ID пожертвования |
| FindDonationById | [FindDonationByIdRequest](#service-FindDonationByIdRequest) | [CreateDonationsResponse](#service-CreateDonationsResponse) |  Поиск пожертвования по ID |
| DeleteDonationByModel | [DeleteDonationByModelRequest](#service-DeleteDonationByModelRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаление пожертвования по модели (например полученного с помощью FindDonationById) |
| DeleteDonationById | [DeleteDonationByIdRequest](#service-DeleteDonationByIdRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаление пожертвование по ID |
| UpdateDonation | [UpdateDonationsRequest](#service-UpdateDonationsRequest) | [CreateDonationsResponse](#service-CreateDonationsResponse) |  Обновление пожертвования |
| Wards | [Empty](#service-Empty) | [WardsResponse](#service-WardsResponse) |  Список всех подопечных |
| CreateWard | [CreateWardRequest](#service-CreateWardRequest) | [Ward](#service-Ward) |  Создание подопечного |
| FindWardById | [FindWardByIdRequest](#service-FindWardByIdRequest) | [Ward](#service-Ward) |  Поиск подопечного по его ID |
| DeleteWardByModel | [Ward](#service-Ward) | [HTTPCodes](#service-HTTPCodes) |  Удаление подопечного по модели (например полученной с помощью FindWardById) |
| DeleteWardById | [DeleteWardByIdRequest](#service-DeleteWardByIdRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаление подопечного по его ID |
| UpdateWard | [Ward](#service-Ward) | [Ward](#service-Ward) |  Обновление подопечного |
| CreateSessions | [CreateSessionRequest](#service-CreateSessionRequest) | [CreateSessionResponse](#service-CreateSessionResponse) |  Создание новой сессии |
| Sessions | [Empty](#service-Empty) | [SessionsResponse](#service-SessionsResponse) |  Список всех сессии |
| FindSessionsById | [FindSessionsByIdRequest](#service-FindSessionsByIdRequest) | [FindSessionsByIdResponse](#service-FindSessionsByIdResponse) |  Поиск сессии по его ID |
| FindSessionsByUserId | [FindSessionsByUserIdRequest](#service-FindSessionsByUserIdRequest) | [FindSessionsByUserIdResponse](#service-FindSessionsByUserIdResponse) |  Поиск сессии по UserID |
| ChangeRefreshTokenById | [ChangeRefreshTokenByIdRequest](#service-ChangeRefreshTokenByIdRequest) | [ChangeRefreshTokenByIdResponse](#service-ChangeRefreshTokenByIdResponse) |  Обновление поля RefreshToken по ID сессии |
| ChangeRefreshTokenByUserId | [ChangeRefreshTokenByUserIdRequest](#service-ChangeRefreshTokenByUserIdRequest) | [ChangeRefreshTokenByUserIdResponse](#service-ChangeRefreshTokenByUserIdResponse) |  Обновление поля RefreshToken по UserID пользователя |
| DeleteSessionByModel | [DeleteSessionByModelRequest](#service-DeleteSessionByModelRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаление сессии по полной модели (например полученной с помощью FindSessionById) |
| DeleteSessionById | [DeleteSessionByIdRequest](#service-DeleteSessionByIdRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаляет сессию по ее ID |
| DeleteSessionByUserId | [DeleteSessionByUserIdRequest](#service-DeleteSessionByUserIdRequest) | [HTTPCodes](#service-HTTPCodes) |  Удаляет сессию по ее UserID |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

