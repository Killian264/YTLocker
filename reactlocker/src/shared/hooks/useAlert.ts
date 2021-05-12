import { useState } from "react";
import { AlertProps } from "../../components/Alert";

export const useAlert = (
	initial: AlertProps
): [AlertProps, (props: AlertProps) => void] => {
	const [alert, _setAlert] = useState(initial);

	const setAlert = (props: AlertProps) => {
		_setAlert(props);
		setTimeout(() => {
			_setAlert(props);
		}, 5000);
	};

	return [alert, setAlert];
};
