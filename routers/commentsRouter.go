package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

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
            Method: "GetCorporatives",
            Router: `/get-corporatives`,
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
            Method: "ResetPin",
            Router: `/reset-pin`,
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

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
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
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
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
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
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
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"] = append(beego.GlobalControllerRouter["msys_payment_app_gateway/controllers:Auth_requestsController"],
        beego.ControllerComments{
            Method: "RegisterAccount",
            Router: `/register-account`,
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

}
