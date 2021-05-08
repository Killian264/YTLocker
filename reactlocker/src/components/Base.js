import { Alert } from "./Alert";
import { LoginPage } from "./LoginPage";
import { LogoBar } from "./LogoBar"

function Base() {
	return (
	<>
		<LogoBar className="absolute top-1 left-5" ></LogoBar>
		<LoginPage/>
	</>
	);
}

export default Base;
