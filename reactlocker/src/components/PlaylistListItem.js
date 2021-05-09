import React from 'react';
import PropTypes from 'prop-types';
import { PlusButton } from './PlusButton';
import no_thumbnail from '../static/no_thumbnail.png';
import { Link } from './Link';


export const PlaylistListItem = ({ className, children, ...props }) => {
	return (
		<div className={`${className} font-bold text-xl hover:bg-secondary-hover rounded-md flex justify-between`}>
			<div className="flex">
				<img src={no_thumbnail} alt="Logo" width="120" className="rounded-lg" />
				<div className="pl-3 flex flex-col">
					<span>Music Memes</span>
					<Link
						className="text-accent"
						href="https://www.youtube.com/playlist?list=PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr"
						target="_blank"
					>Youtube Playlist</Link>
				</div>
			</div>
			<div className="mr-1 text-3xl my-auto select-none">
				>
			</div>
		</div>
	);
};

PlaylistListItem.propTypes = {};
