import React from 'react';


export const Link = ({ className, children, ...props }) => {
	return (
		<a className={`${className} text-sm underline text-secondary-text cursor-pointer select-none`} {...props} >{children}</a>
	);
};