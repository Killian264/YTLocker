import React from "react";

export interface LinkProps extends React.AnchorHTMLAttributes<HTMLAnchorElement> {
	className?: string;
	pointer?: boolean;
}

export const Link: React.FC<LinkProps> = ({ className, children, pointer = true, ...props }) => {
	let pointerCSS = pointer ? "cursor-pointer" : "cursor-default";
	let css = `${className} ${pointerCSS} text-sm text-secondary-text select-none hover:opacity-75 hover:underline`;

	return (
		<a className={css} {...props}>
			{children}
		</a>
	);
};
