import axios from "axios";
import { useEffect } from "react";
import { BrowserRouter as Router, Redirect, Route, RouteProps, Switch, useHistory } from "react-router-dom";
import { DROPLET_BASE } from "../shared/env";
import { useBearer } from "../shared/hooks/useBearer";
import { DashboardPage } from "./DashboardPage";
import { LoginPage } from "./LoginPage";

axios.defaults.baseURL = DROPLET_BASE;

const AuthenticatedRoute: React.FC<RouteProps> = ({ ...props }) => {
	const [bearer] = useBearer("");
	const history = useHistory();

	useEffect(() => {
		axios.defaults.headers.common["Authorization"] = bearer;

		let id = axios.interceptors.response.use(
			(response) => response,
			(error) => {
				if (typeof error.response != "undefined" && error.response.status === 401) {
					history.push("/login");
				}
				return error;
			}
		);
		return () => {
			axios.interceptors.request.eject(id);
		};
	});

	if (bearer === "") {
		return <Redirect to="/login" />;
	}

	return <Route {...props}></Route>;
};

export const IndexPage: React.FC<{}> = () => {
	return (
		<Router>
			<Switch>
				<AuthenticatedRoute path="/" exact component={DashboardPage}></AuthenticatedRoute>
				<Route path="/login" exact component={LoginPage}></Route>
			</Switch>
		</Router>
	);
};
