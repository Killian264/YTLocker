import React from 'react';
import PropTypes from 'prop-types';
import { PlusButton } from './PlusButton';
import no_thumbnail from '../static/no_thumbnail.png';


const sizes = {
	large: "py-2 px-6 text-sm rounded-lg",
	medium: "py-1.5 px-5 text-sm rounded-md",
	small: "py-1 px-3 text-xs rounded-md",
}

const colors = {
	primary:   "text-accent-text    bg-accent    hover:bg-accent-hover    disabled:bg-accent-disabled    disabled:text-accent-text-disabled",
	secondary: "text-secondary-text bg-secondary hover:bg-secondary-hover disabled:bg-secondary-disabled disabled:text-secondary-text-disabled"
}


export const PlaylistListItem = ({ children, ...props }) => {
	return (
		<div className={`bg-secondary py-4 px-6 rounded-md`}>
			<div className="flex justify-between">
				<div className="text-2xl">
					<span className="font-bold inline-block align-middle leading-none">Playlists</span>
				</div>
				<PlusButton></PlusButton>
			</div>
			<div className="border-b-2 mt-2 mb-3"></div>
			<div className="font-bold text-xl">
				<div className="flex">
					<img src={no_thumbnail} alt="Logo" width="120" className="rounded-lg" />
					<div className="pl-3">
						<span>Music Memes</span>
					</div>
				</div>
			</div>
		</div>
	);
};

PlaylistListItem.propTypes = {};
