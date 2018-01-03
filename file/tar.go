package file

// func Tar() {
// 	buf := &bytes.Buffer{}
// 	tw := tar.NewWriter(buf)
// 	defer tw.Close()

// 	for _, dir := range []string{".", "conf.d"} {
// 		hdr := &tar.Header{
// 			Name:       dir,
// 			Mode:       0755,
// 			Uid:        os.Getuid(),
// 			Gid:        os.Getgid(),
// 			ModTime:    time.Now(),
// 			Typeflag:   tar.TypeDir,
// 			AccessTime: time.Now(),
// 			ChangeTime: time.Now(),
// 		}
// 		err := tw.WriteHeader(hdr)
// 		if err != nil {
// 			log.Debugln("WriteHeader,", err)
// 			return nil, err
// 		}
// 	}
// 	for _, file := range filelist {
// 		hdr := &tar.Header{
// 			Name:       file.Path,
// 			Mode:       int64(file.Perm),
// 			Uid:        os.Getuid(),
// 			Gid:        os.Getgid(),
// 			Size:       int64(len(file.Content)),
// 			ModTime:    time.Now(),
// 			Typeflag:   tar.TypeReg,
// 			AccessTime: time.Now(),
// 			ChangeTime: time.Now(),
// 		}
// 		err := tw.WriteHeader(hdr)
// 		if err != nil {
// 			log.Debugln("WriteHeader,", err)
// 			return nil, err
// 		}
// 		n, err := tw.Write(file.Content)
// 		if err != nil {
// 			log.Debugln("Write,", err, n)
// 			return nil, err
// 		}
// 		tw.Flush()
// 	}

// 	bs := buf.Bytes()

// 	fullpath := "nginx.tar"
// 	filename, _, typ := parseFilename(fullpath)

// 	sha1, size := SHA1(bs)
// 	return &nevis_common.File{
// 		Name:    filename,
// 		Path:    filepath.Join(rootdir, filename),
// 		Type:    typ,
// 		Perm:    0644,
// 		Size_:   size,
// 		Sha1:    sha1,
// 		Content: bs,
// 	}, nil
// }
