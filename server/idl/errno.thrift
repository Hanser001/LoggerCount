namespace go errno

enum Err {
    Success              = 0,
    ParamsErr            = 1,
    ServiceErr           = 2,
    RPCUserErr           = 10000,
    UserServerErr        = 10001,
    UserAlreadyExistErr  = 10002,
    UserNotFoundErr      = 10003,
    AuthorizeFailErr     = 10004,
    RPCFileErr           = 20000,
    FileServerErr        = 20001,
}