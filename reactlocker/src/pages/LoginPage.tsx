import { RouteComponentProps } from "react-router-dom";
import { GoogleOAuthCard } from "../components/GoogleOauthCard";
import { LogoBar } from "../components/LogoBar";


export interface LoginPageProps extends RouteComponentProps {
	className?: string;
}

export const LoginPage: React.FC<LoginPageProps> = ({ className, history }) => {
	return (
		<>
			<LogoBar className="absolute top-1 left-5"></LogoBar>
			<div className="flex h-screen">
				<div className="m-auto">
					<GoogleOAuthCard
						type="login"
					/>
				</div>
			</div>
		</>
	);
};

LoginPage.propTypes = {};
