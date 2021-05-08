import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import image from '../static/logo.png'

export const LogoBar = ({ className, children, ...props }) => {

	let pos = "inset-x-0 mx-auto mt-3"
	let css = `flex ${pos} ${className}`

	return (
		<div
			type="button"
			className={css}
			{...props}
		>
			<img src={image} alt="Logo" width="32" height="32"  />
			<span className="text-xl font-bold ml-1 text-accent tracking-wide" >YTLocker</span>
		</div>
	);
};

LogoBar.propTypes = {
	className: PropTypes.string,
};
