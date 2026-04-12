package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "AccountBalance",
            Router: `/account-balance`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "ListAccountDetails",
            Router: `/account-details`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "Deposit",
            Router: `/deposit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "GetAgentTransactions",
            Router: `/get-agent-transactions`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "GetCorporatives",
            Router: `/get-corporatives`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "GetBilTransactionWithTransactionRef",
            Router: `/get-transaction-by-reference`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "GetUserDetails",
            Router: `/get-user-details`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "ListAccountLoans",
            Router: `/list-account-loans`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Agent_api_requestsController"],
        beego.ControllerComments{
            Method: "LoanRepayment",
            Router: `/loan-repayment`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "AccountBalance",
            Router: `/account-balance`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "AccountQuery",
            Router: `/account-query`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "BuyAirtime",
            Router: `/buy-airtime`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "BuyDataBundle",
            Router: `/buy-data-bundle`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "Deposit",
            Router: `/deposit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GetBilTransactionWithTransactionRef",
            Router: `/get-biller-transaction-by-reference`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GetBilTransactions",
            Router: `/get-biller-transaction-history`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GetCorporatives",
            Router: `/get-corporatives`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GetCustomerAccountHistory",
            Router: `/get-customer-account-history`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GetCustomerAccountStatement",
            Router: `/get-customer-account-statement`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GetCustomerAccounts",
            Router: `/get-customer-accounts`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GetCustomerDetails",
            Router: `/get-customer-details`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GetBundles",
            Router: `/get-data-bundles`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GetPaymentMethods",
            Router: `/get-payment-methods`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "GhanaWaterAccountQuery",
            Router: `/ghana-water-account-query`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "NameInquiry",
            Router: `/name-inquiry`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "PayDSTV",
            Router: `/pay-dstv`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "PayECG",
            Router: `/pay-ecg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "PayGOTV",
            Router: `/pay-gotv`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "PayStartimesTvBill",
            Router: `/pay-startimes`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "PayWaterBill",
            Router: `/pay-water`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "RegisterAccount",
            Router: `/register-account`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "ResetPin",
            Router: `/reset-pin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "CloseAccount",
            Router: `/v2/close-account`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "ValidateCustomer",
            Router: `/validate-customer`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Api_requestsController"],
        beego.ControllerComments{
            Method: "Withdrawal",
            Router: `/withdrawal`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "ChangePassword",
            Router: `/change-password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "RefreshAccessToken",
            Router: `/refresh-customer-access-token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "RefreshUserAccessToken",
            Router: `/refresh-user-access-token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "RegisterUser",
            Router: `/register-user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: `/reset-password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "LoginUser",
            Router: `/user/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:CallbackController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:CallbackController"],
        beego.ControllerComments{
            Method: "Callback",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:CallbackController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:CallbackController"],
        beego.ControllerComments{
            Method: "CheckTransactionStatus",
            Router: `/check-transaction-status`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:CallbackController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:CallbackController"],
        beego.ControllerComments{
            Method: "RequestMoneyCallback",
            Router: `/payment-callback`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:CallbackController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:CallbackController"],
        beego.ControllerComments{
            Method: "TransferCallback",
            Router: `/transfer-callback`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
