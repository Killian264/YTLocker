import axios from "axios";
import { useEffect, useState } from "react";
import { RouteComponentProps } from "react-router-dom";
import { GoogleOAuthCard } from "../components/GoogleOauthCard";
import { LogoBar } from "../components/LogoBar";
import { useBearer } from "../hooks/useBearer";

export interface LoginPageProps extends RouteComponentProps {
	className?: string;
}

export const LoginPage: React.FC<LoginPageProps> = ({ className, history }) => {
	const [isLoading, setLoading] = useState(true);
	const [bearer, setBearer] = useBearer("");

	useEffect(() => {
		axios.post("/user/session/create", {}, {}).then((response) => {
			setBearer(response.data.Data.Bearer);
			setLoading(false);
		});
	}, [setBearer]);

	return (
		<>
			<LogoBar className="absolute top-1 left-5"></LogoBar>
			<div className="flex h-screen">
				<div className="m-auto">
					<GoogleOAuthCard type="login" loading={isLoading} bearer={bearer} />
				</div>
			</div>
		</>
	);
};

LoginPage.propTypes = {};
