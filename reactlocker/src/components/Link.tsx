import React from "react";

export interface LinkProps
	extends React.AnchorHTMLAttributes<HTMLAnchorElement> {
	className?: string;
}

export const Link: React.FC<LinkProps> = ({
	className,
	children,
	...props
}) => {
	let css = `${className} text-sm underline text-secondary-text cursor-pointer select-none`;

	return (
		<a className={css} {...props}>
			{children}
		</a>
	);
};
