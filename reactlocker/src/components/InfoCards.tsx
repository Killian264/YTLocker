import React from "react";

interface ClassNameProp {
	className?: string;
}

export const BaseAccountInfo: React.FC<ClassNameProp> = ({ className }) => {
	return (
		<div className={`${className} bg-primary-700 p-10 pt-8 rounded-md`}>
			<span className="text-2xl font-bold">YTLocker Playlists</span>
			<div className="ml-1 mt-1">
				The YTLocker Playlists account can be used by any user. To ensure privacy, playlists are
				unlisted, meaning they can only be found with a link.
			</div>
			<div className="mt-4">
				<span className="text-lg font-bold">Account Limitations</span>
				<ul className="mt-1 ml-5 list-disc">
					<li>
						<span className="font-semibold">Shared</span> - The account can be used to create
						playlists by any user.
					</li>
					<li>
						<span className="font-semibold">Slower</span> - Videos may take longer to be added to
						the playlist.
					</li>
					<li>
						<span className="font-semibold">Limited</span> - Number of playlists on this account
						is limited.
					</li>
				</ul>
			</div>
		</div>
	);
};

export const BeforeLinkingInfo: React.FC<ClassNameProp> = ({ className }) => {
	return (
		<div className={`${className} bg-primary-700 p-10 pt-8 rounded-md`}>
			<span className="text-2xl font-bold">Account Linking</span>
			<div className="ml-1 mt-1">
				To create playlists on a specific account, it must be linked with management permissions.
				Because of the permission level, YTLocker also gains the permissions listed below. If you're
				worried about privacy or security, creating and linking a secondary Youtube account is
				recommended.
			</div>
			<div className="mt-4">
				<span className="text-lg font-bold">Manage your Youtube account</span>
				<ul className="mt-1 ml-5 list-disc">
					<li>View and manage your videos and playlists</li>
					<li>View and manage your YouTube activity, including posting public comments</li>
				</ul>
			</div>
		</div>
	);
};

export const UsingYTLockerInfo: React.FC<ClassNameProp> = ({ className }) => {
	return (
		<div className={`${className} bg-primary-700 p-10 pt-8 rounded-md`}>
			<span className="text-2xl font-bold">Using YTLocker</span>
			<div className="ml-1 mt-1">
				YTLocker is a platform for management of Youtube playlists. It allows you to create a playlist
				and have videos be automatically added within 2 minutes of upload. The primary use case is for
				automatically adding music content to a playlist, these steps to create a music playlist are
				listed below
			</div>
			<div className="mt-4">
				<span className="text-lg font-bold">Creating a Music Playlist</span>
				<ul className="mt-1 ml-5 list-disc">
					<li>Create a playlist using the YTLocker account or link an account.</li>
					<li>
						Add channels to your playlist, uploads from these channels will be added to the
						playlist.
					</li>
					<li>Bookmark your new playlist and wait.</li>
				</ul>
			</div>
		</div>
	);
};

export const NeedHelpInfo: React.FC<ClassNameProp> = ({ className }) => {
	return (
		<div className={`${className} bg-primary-700 p-10 pt-8 rounded-md`}>
			<span className="text-2xl font-bold">Found an Issue?</span>
			<div className="ml-1 mt-1">
				YTLocker is open-source. If you've found a bug please click Issues at the bottom of the page
				and create an issue.
			</div>
		</div>
	);
};
