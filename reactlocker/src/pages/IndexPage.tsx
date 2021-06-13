import axios from "axios";
import { BrowserRouter as Router, Redirect, Route, RouteProps, Switch, useHistory } from "react-router-dom";
import { DROPLET_BASE } from "../shared/env";
import { useBearer } from "../shared/hooks/useBearer";
import { DashboardPage } from "./DashboardPage";
import { LoginPage } from "./LoginPage";

const AuthenticatedRoute: React.FC<RouteProps> = ({ ...props }) => {
	const [bearer] = useBearer("");

	if (bearer === "") {
		return <Redirect to="/login" />;
	}

	return <Route {...props}></Route>;
};

const AxiosInteceptorProvider: React.FC<{}> = ({ children }) => {
	const history = useHistory();

	axios.defaults.baseURL = DROPLET_BASE;

	axios.interceptors.response.use(
		function (response) {
			return response;
		},
		function (error) {
			if (error.response && error.response.status === 401) {
				history.push("/login");
			}
			return error;
		}
	);

	return <> {children} </>;
};

export const IndexPage: React.FC<{}> = () => {
	return (
		<AxiosInteceptorProvider>
			<Router>
				<Switch>
					<AuthenticatedRoute path="/" exact component={DashboardPage}></AuthenticatedRoute>
					<Route path="/login" exact component={LoginPage}></Route>
				</Switch>
			</Router>
		</AxiosInteceptorProvider>
	);
};
