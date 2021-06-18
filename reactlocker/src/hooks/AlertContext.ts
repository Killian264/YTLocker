import { createContext } from "react";
import { AlertProps } from "../components/Alert";

export const AlertContext = createContext<{ pushAlert: (alert: AlertProps) => void }>({
	pushAlert: (alert) => {},
});
