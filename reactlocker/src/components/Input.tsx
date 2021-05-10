import React from "react";

export interface InputProps extends React.HTMLAttributes<HTMLInputElement> {
	className?: string;
	value: string;
	type?: string;
}

export const Input: React.FC<InputProps> = ({
	className,
	value,
	type = "",
	...props
}) => {
	const bg = `bg-primary-700`;
	const ring = `focus:outline-none`;
	const c = `${className} ${ring} ${bg} w-full py-2 px-4 rounded-md text-primary-100 placeholder-primary-300 focus:text-secondary-text`;

	return <input className={c} type={type} value={value} {...props} />;
};
